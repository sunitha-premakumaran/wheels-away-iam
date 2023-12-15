package services

import (
	"context"

	"github.com/sunitha/wheels-away-iam/internal/core/domain"
)

type (
	UserRepository interface {
		GetUsers(ctx context.Context) ([]*domain.DecoratedUser, error)
	}
)
