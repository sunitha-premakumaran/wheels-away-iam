package services

import (
	"context"

	"github.com/sunitha/wheels-away-iam/internal/core/domain"
)

type (
	UserRepository interface {
		SaveUser(ctx context.Context, user *domain.User) error
		GetUsers(ctx context.Context, page, size int,
			searchKey *domain.UserSearhKey, searchString *string) (
			[]*domain.DecoratedUser, *domain.PageInfo, error)
	}

	RoleRepository interface {
		GetRolesByIDs(ctx context.Context, roleIDs []string) ([]*domain.Role, error)
		GetRoles(ctx context.Context) ([]*domain.Role, error)
		SaveRole(ctx context.Context, role *domain.Role) error
	}

	RoleUserMappingRepository interface {
		SaveRoleUserMapping(ctx context.Context, roleUserMap *domain.RoleUserMapping) error
	}
)
