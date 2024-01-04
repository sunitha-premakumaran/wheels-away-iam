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

func (i *UserInteractor) GetUsers(ctx context.Context, page, size int,
	searchKey *domain.UserSearhKey, searchString *string) ([]*domain.DecoratedUser, *domain.PageInfo, error) {
	return i.userRepo.GetUsers(ctx, page, size, searchKey, searchString)
}

func (i *UserInteractor) GetUser(ctx context.Context, userID string) (*domain.DecoratedUser, error) {
	return i.userRepo.GetUser(ctx, userID)
}

func (i *UserInteractor) SaveUser(ctx context.Context, user *domain.User) error {
	return i.userRepo.SaveUser(ctx, user)
}
