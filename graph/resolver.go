package graph

import (
	"context"

	"github.com/sunitha/wheels-away-iam/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type UserProcessor interface {
	GetUsers(ctx context.Context) ([]*model.User, error)
}

type Resolver struct {
	UserProcessor UserProcessor
}