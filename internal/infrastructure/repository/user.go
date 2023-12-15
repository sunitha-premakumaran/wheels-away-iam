package repository

import (
	"context"
	"errors"

	"github.com/sunitha/wheels-away-iam/internal/core/domain"
	"github.com/sunitha/wheels-away-iam/internal/infrastructure/tables"
	"gorm.io/gorm"
)

type UserRepository struct {
	gormDB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		gormDB: db,
	}
}

func (r *UserRepository) GetUsers(ctx context.Context) ([]*domain.DecoratedUser, error) {
	var users []*tables.User
	result := r.gormDB.Model(&tables.User{}).Find(&users)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	du := make([]*domain.DecoratedUser, 0, len(users))
	for _, u := range users {
		us := &domain.DecoratedUser{
			User: u.ToUserDomain(),
		}
		// for _, r := range u.Roles {
		// 	us.UserRoles = append(us.UserRoles, r.ToRoleDomain())
		// }
		du = append(du, us)
	}

	return du, nil
}
