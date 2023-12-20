package graph

import (
	"context"

	"github.com/sunitha/wheels-away-iam/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type UserProcessor interface {
	CreateUser(ctx context.Context, user *model.UserInput) (*model.UpsertResponse, error)
	GetUsers(ctx context.Context, pageInput model.PageInput, searchInput *model.UserSearchInput) (*model.UserResponse, error)
}

type RoleProcessor interface {
	GetRoles(ctx context.Context) ([]*model.Role, error)
	CreateRole(ctx context.Context, role *model.RoleInput) (*model.UpsertResponse, error)
}

type Resolver struct {
	UserProcessor UserProcessor
	RoleProcessor RoleProcessor
}
