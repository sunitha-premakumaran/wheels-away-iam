package repository

import (
	"context"

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
		CreatedBy:     role.CreatedBy,
		CreatedAt:     role.CreatedAt,
		LastUpdatedAt: role.LastUpdatedAt,
		LastUpdatedBy: role.LastUpdatedBy,
		DeletedAt:     role.DeletedAt,
		DeletedBy:     role.DeletedBy,
	}
}
