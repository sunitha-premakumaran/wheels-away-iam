package services

import (
	"context"

	"github.com/sunitha/wheels-away-iam/internal/core/domain"
)

type UserInteractor struct {
	userRepo UserRepository
}

func NewUserInteractor(userRepo UserRepository) *UserInteractor {
	return &UserInteractor{
		userRepo: userRepo,
	}
}

func (i *UserInteractor) GetUsers(ctx context.Context) ([]*domain.DecoratedUser, error) {
	return i.userRepo.GetUsers(ctx)
}
