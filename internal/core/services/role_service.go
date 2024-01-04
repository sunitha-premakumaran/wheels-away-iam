package services

import (
	"context"

	"github.com/sunitha/wheels-away-iam/internal/core/domain"
)

type RoleInteractor struct {
	roleRepo RoleRepository
}

func NewRoleInteractor(roleRepo RoleRepository) *RoleInteractor {
	return &RoleInteractor{
		roleRepo: roleRepo,
	}
}

func (i *RoleInteractor) SaveRole(ctx context.Context, role *domain.Role) error {
	return i.roleRepo.SaveRole(ctx, role)
}

func (i *RoleInteractor) GetRolesByIDs(ctx context.Context, roleIDs []string) ([]*domain.Role, error) {
	return i.roleRepo.GetRolesByIDs(ctx, roleIDs)
}

func (i *RoleInteractor) GetRole(ctx context.Context, roleID string) (*domain.Role, error) {
	return i.roleRepo.GetRole(ctx, roleID)
}

func (i *RoleInteractor) GetRoles(ctx context.Context) ([]*domain.Role, error) {
	return i.roleRepo.GetRoles(ctx)
}
