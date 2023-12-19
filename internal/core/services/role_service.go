package services

import (
	"context"

	"github.com/sunitha/wheels-away-iam/internal/core/domain"
)

type RoleInteractor struct {
	userRepo RoleRepository
}

func NewRoleInteractor(userRepo RoleRepository) *RoleInteractor {
	return &RoleInteractor{
		userRepo: userRepo,
	}
}

func (i *RoleInteractor) GetRoles(ctx context.Context) ([]*domain.Role, error) {
	return i.userRepo.GetRoles(ctx)
}

func (i *RoleInteractor) SaveRole(ctx context.Context, user *domain.Role) error {
	return i.userRepo.SaveRole(ctx, user)
}
