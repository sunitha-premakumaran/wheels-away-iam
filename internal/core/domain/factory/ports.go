package factory

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
		GetRolesByIDs(ctx context.Context, roleIDs []string) ([]*domain.Role, error)
	}

	UserIDPInteractor interface {
		CreateIDPUser(ctx context.Context, user *domain.User) (string, error)
		CreateUserGrant(ctx context.Context, userID string, roles []string) error
	}

	RoleIDPInteractor interface {
		SaveIDPRole(ctx context.Context, role *domain.Role) (string, error)
	}

	RoleUserMappingInteractor interface {
		SaveRoleUserMapping(ctx context.Context,
			roleUserMapping *domain.RoleUserMapping) error
	}
)
