package services

import (
	"context"

	"github.com/sunitha/wheels-away-iam/internal/core/domain"
)

type RoleUserMappingInteractor struct {
	roleUserMappingRepo RoleUserMappingRepository
}

func NewRoleUserMappingInteractor(
	roleUserMappingRepo RoleUserMappingRepository) *RoleUserMappingInteractor {
	return &RoleUserMappingInteractor{
		roleUserMappingRepo: roleUserMappingRepo,
	}
}

func (i *RoleUserMappingInteractor) SaveRoleUserMapping(ctx context.Context,
	roleUserMapping *domain.RoleUserMapping) error {
	return i.roleUserMappingRepo.SaveRoleUserMapping(ctx, roleUserMapping)
}
