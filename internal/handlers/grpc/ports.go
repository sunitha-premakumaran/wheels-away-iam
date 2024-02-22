package grpc

import (
	"context"

	"github.com/sunitha/wheels-away-iam/internal/core/domain"
)

type (
	UserInteractor interface {
		GetUser(ctx context.Context, userID string) (*domain.DecoratedUser, error)
	}
)
