package api

import (
	"context"

	"github.com/sunitha/wheels-away-iam/internal/core/domain"
)

type (
	UserInteractor interface {
		GetUsers(ctx context.Context, page, size int,
			searchKey *domain.UserSearhKey, searchString *string) (
			[]*domain.DecoratedUser, *domain.PageInfo, error)
		SaveUser(ctx context.Context, user *domain.User) error
	}

	RoleInteractor interface {
		GetRoles(ctx context.Context) ([]*domain.Role, error)
		SaveRole(ctx context.Context, role *domain.Role) error
	}
)
