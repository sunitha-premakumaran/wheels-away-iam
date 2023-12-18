package worker

import (
	"context"

	"github.com/sunitha/wheels-away-iam/internal/core/domain"
)

var (
	UserInteractor interface {
		GetUsers(ctx context.Context) ([]*domain.DecoratedUser, error)
		SaveUser(ctx context.Context, user *domain.User) error
	}

	RoleInteractor interface {
		GetRoles(ctx context.Context) ([]*domain.Role, error)
		SaveRole(ctx context.Context, role *domain.Role) error
	}

	UserIDPInteractor interface {
		CreateIDPUser(ctx context.Context, user *domain.User) (string, error)
	}

	RoleIDPInteractor interface {
		CreateIDPRole(ctx context.Context, role *domain.Role) error
	}
)
