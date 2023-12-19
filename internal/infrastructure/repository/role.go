package repository

import (
	"context"
	"errors"

	"github.com/lib/pq"
	"github.com/sunitha/wheels-away-iam/internal/core/domain"
	"github.com/sunitha/wheels-away-iam/internal/infrastructure/models/tables"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RoleRepository struct {
	gormDB *gorm.DB
}

func NewRoleRepository(gormDB *gorm.DB) *RoleRepository {
	return &RoleRepository{
		gormDB: gormDB,
	}
}

func (r *RoleRepository) GetRolesByIDs(ctx context.Context, roleIDs []string) ([]*domain.Role, error) {
	return r.getRolesByIDs(ctx, roleIDs)
}

func (r *RoleRepository) getRolesByIDs(ctx context.Context, roleIDs []string) ([]*domain.Role, error) {
	var roles []*tables.Role
	result := r.gormDB.Model(&tables.Role{}).Where("role_pk IN (?)", pq.StringArray(roleIDs)).Find(&roles)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	rr := make([]*domain.Role, 0, len(roles))
	for _, r := range roles {
		rr = append(rr, r.ToRoleDomain())
	}
	return rr, nil
}

func (r *RoleRepository) SaveRole(ctx context.Context, role *domain.Role) error {
	roleToUpsert := mapRoleDomainToTable(role)
	return r.saveRole(ctx, roleToUpsert)
}

func (r *RoleRepository) saveRole(ctx context.Context, role *tables.Role) error {
	tx := r.gormDB.WithContext(ctx)
	result := tx.Session(&gorm.Session{FullSaveAssociations: true}).Clauses(clause.OnConflict{
		Where: clause.Where{Exprs: []clause.Expression{
			clause.Lt{
				Column: `"roles"."lastupdated_at"`,
				Value:  role.LastUpdatedAt,
			},
			clause.Or(
				clause.Eq{
					Column: `"roles"."lastupdated_at"`,
					Value:  nil,
				},
			),
		}},
		UpdateAll: true,
	}).Create(&role)

	return result.Error
}

func mapRoleDomainToTable(role *domain.Role) *tables.Role {
	scopes := []string{}
	for _, e := range role.Scopes {
		scopes = append(scopes, string(e))
	}
	return &tables.Role{
		UUID:          role.UUID,
		Name:          role.Name,
		Description:   role.Description,
		Scopes:        scopes,
		AuthKey:       role.AuthKey,
		CreatedBy:     role.CreatedBy,
		CreatedAt:     role.CreatedAt,
		LastUpdatedAt: role.LastUpdatedAt,
		LastUpdatedBy: role.LastUpdatedBy,
		DeletedAt:     role.DeletedAt,
		DeletedBy:     role.DeletedBy,
	}
}
