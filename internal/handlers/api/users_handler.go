package api

import (
	"context"
	"fmt"

	"github.com/sunitha/wheels-away-iam/graph/model"
	"github.com/sunitha/wheels-away-iam/internal/core/domain"
	"github.com/sunitha/wheels-away-iam/internal/core/domain/factory"
	"github.com/sunitha/wheels-away-iam/internal/core/enums"
)

var (
	userSearchKey = map[model.UserSearchKey]domain.UserSearhKey{
		model.UserSearchKeyEmail: domain.Email,
		model.UserSearchKeyName:  domain.Name,
	}
)

type UserProcessor struct {
	userInteractor            UserInteractor
	userIDPInteractor         UserIDPInteractor
	roleUserMappingInteractor RoleUserMappingInteractor
	roleInteractor            RoleInteractor
}

func NewUserProcessor(userInteractor UserInteractor, userIDPInteractor UserIDPInteractor,
	roleUserMappingInteractor RoleUserMappingInteractor,
	roleInteractor RoleInteractor) *UserProcessor {
	return &UserProcessor{
		userIDPInteractor:         userIDPInteractor,
		userInteractor:            userInteractor,
		roleUserMappingInteractor: roleUserMappingInteractor,
		roleInteractor:            roleInteractor,
	}
}

func (p *UserProcessor) CreateUser(ctx context.Context, user *model.UserInput) (*model.UpsertResponse, error) {
	uFactory := factory.NewCreateUserFactory(p.userInteractor, p.userIDPInteractor, p.roleInteractor, p.roleUserMappingInteractor)
	de, err := uFactory.Create(ctx, user.FirstName, user.LastName, user.Email, user.Phone, nil, nil, enums.INACTIVE, user.UserRoles, domain.SystemUUID)
	if err != nil {
		return &model.UpsertResponse{
			Success: false,
			ErrorMessage: &model.ErrorMessage{
				Code: string(de),
				Msg:  "",
			},
		}, nil
	}
	return &model.UpsertResponse{
		Success:      true,
		ErrorMessage: nil,
	}, nil
}

func (p *UserProcessor) GetUsers(ctx context.Context, pageInput model.PageInput, searchInput *model.UserSearchInput) (*model.UserResponse, error) {
	var searckKey domain.UserSearhKey
	var searchString string
	if searchInput != nil && searchInput.SearchString != "" {
		var ok bool
		searckKey, ok = userSearchKey[searchInput.SearchKey]
		if !ok {
			return nil, fmt.Errorf("User search key is not valid: %s", searchInput.SearchKey)
		}
		searchString = searchInput.SearchString
	}
	du, pageInfo, err := p.userInteractor.GetUsers(ctx, pageInput.PageNumber, pageInput.PageSize, &searckKey, &searchString)
	if err != nil {
		return nil, fmt.Errorf("error getting users: %w", err)
	}
	return &model.UserResponse{
		Users:    p.mapDecoratedUsersToModel(du),
		PageInfo: p.mapPagedInfoToModel(pageInfo),
	}, nil
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

func (p *UserProcessor) mapPagedInfoToModel(pageInfo *domain.PageInfo) *model.PageInfo {
	return &model.PageInfo{
		PageSize:   pageInfo.PageSize,
		PageNumber: pageInfo.CurrentPage,
		TotalItems: pageInfo.TotalItems,
		TotalPages: pageInfo.TotalPages,
	}
}

func statusToModel(status enums.UserStatus) model.UserStatus {
	if status == enums.ACTIVE {
		return model.UserStatusActive
	} else if status == enums.INACTIVE {
		return model.UserStatusInActive
	}
	return ""
}
