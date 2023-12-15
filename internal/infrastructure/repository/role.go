package repository

import "github.com/sunitha/wheels-away-iam/pkg/gorm"

type RoleRepository struct {
	db *gorm.DBClient
}

func NewRoleRepository(db *gorm.DBClient) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}
