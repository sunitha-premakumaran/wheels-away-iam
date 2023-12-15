package api

import (
	"context"

	"github.com/sunitha/wheels-away-iam/internal/core/domain"
)

type (
	UserInteractor interface {
		GetUsers(ctx context.Context) ([]*domain.DecoratedUser, error)
	}
)
