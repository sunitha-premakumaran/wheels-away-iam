package api

import (
	"context"
	"fmt"

	"github.com/sunitha/wheels-away-iam/graph/model"
	"github.com/sunitha/wheels-away-iam/internal/core/domain"
	"github.com/sunitha/wheels-away-iam/internal/core/enums"
)

type UserProcessor struct {
	userInteractor UserInteractor
}

func NewUserProcessor(userInteractor UserInteractor) *UserProcessor {
	return &UserProcessor{
		userInteractor: userInteractor,
	}
}

func (p *UserProcessor) GetUsers(ctx context.Context) ([]*model.User, error) {
	du, err := p.userInteractor.GetUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting users: %w", err)
	}
	return p.mapDecoratedUsersToModel(du), nil
}

func (p *UserProcessor) mapDecoratedUsersToModel(du []*domain.DecoratedUser) []*model.User {
	mu := make([]*model.User, 0, len(du))
	for _, u := range du {
		var roles []*string
		for _, r := range u.UserRoles {
			roles = append(roles, &r.Name)
		}
		mu = append(mu, &model.User{
			FirstName: u.User.FirstName,
			LastName:  u.User.LastName,
			Email:     u.User.Email,
			Phone:     u.User.Phone,
			Status:    statusToModel(u.User.Status),
			UserRoles: roles,
		})
	}
	return mu
}

func statusToModel(status enums.UserStatus) model.UserStatus {
	if status == enums.ACTIVE {
		return model.UserStatusActive
	} else if status == enums.INACTIVE {
		return model.UserStatusInActive
	}
	return ""
}
