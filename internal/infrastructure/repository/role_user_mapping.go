package repository

import (
	"context"

	"github.com/sunitha/wheels-away-iam/internal/core/domain"
	"github.com/sunitha/wheels-away-iam/internal/infrastructure/models/tables"
	"gorm.io/gorm"
)

type RoleUserMappingRepository struct {
	gormDB *gorm.DB
}

func NewRoleUserMappingRepository(gormDB *gorm.DB) *RoleUserMappingRepository {
	return &RoleUserMappingRepository{
		gormDB: gormDB,
	}
}

func (r *RoleUserMappingRepository) SaveRoleUserMapping(ctx context.Context, roleUserMap *domain.RoleUserMapping) error {
	roleUserMapToUpsert := mapRoleUserMappingDomainToTable(roleUserMap)
	return r.saveRoleUserMapping(ctx, roleUserMapToUpsert)
}

func (r *RoleUserMappingRepository) saveRoleUserMapping(ctx context.Context, roleUserMap *tables.RoleUserMapping) error {
	tx := r.gormDB.WithContext(ctx)
	result := tx.Session(&gorm.Session{FullSaveAssociations: true}).Create(&roleUserMap)
	return result.Error
}

func mapRoleUserMappingDomainToTable(roleUserMap *domain.RoleUserMapping) *tables.RoleUserMapping {
	return &tables.RoleUserMapping{
		UUID:      roleUserMap.UUID,
		RoleID:    roleUserMap.RoleID,
		UserID:    roleUserMap.UserID,
		CreatedBy: roleUserMap.CreatedBy,
		CreatedAt: roleUserMap.CreatedAt,
		DeletedAt: roleUserMap.DeletedAt,
		DeletedBy: roleUserMap.DeletedBy,
	}
}
