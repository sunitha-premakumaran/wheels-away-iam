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
	userScopeMap = map[model.UserPermision]enums.UserScope{
		model.UserPermisionRoleRead:  enums.ROLES_READ,
		model.UserPermisionRoleWrite: enums.ROLES_WRITE,
		model.UserPermisionUserRead:  enums.USERS_READ,
		model.UserPermisionUserWrite: enums.USERS_WRITE,
	}
)

type CreateRoleProcessor struct {
	roleInteractor    RoleInteractor
	roleIDPInteractor RoleIDPInteractor
}

func NewCreateRoleProcessor(
	roleInteractor RoleInteractor,
	roleIDPInteractor RoleIDPInteractor,
) *CreateRoleProcessor {
	return &CreateRoleProcessor{
		roleInteractor:    roleInteractor,
		roleIDPInteractor: roleIDPInteractor,
	}
}

func (h *CreateRoleProcessor) GetRoles(ctx context.Context) ([]*model.Role, error) {
	roles, err := h.roleInteractor.GetRoles(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while fetching row: %w", err)
	}
	var mRoles []*model.Role
	for _, r := range roles {
		mRoles = append(mRoles, mapDomainRoleToModel(r))
	}
	return mRoles, nil
}

func (h *CreateRoleProcessor) CreateRole(ctx context.Context, role *model.RoleInput) (*model.UpsertResponse, error) {
	var scopes []enums.UserScope
	for _, scope := range role.Permissions {
		s, ok := userScopeMap[scope]
		if !ok {
			return &model.UpsertResponse{
				Success: false,
				ErrorMessage: &model.ErrorMessage{
					Code: string(domain.BadRequestError),
					Msg:  fmt.Sprintf("%s is not a valid scope", scope),
				},
			}, nil
		}
		scopes = append(scopes, s)
	}
	rFactory := factory.NewCreateRoleFactory(h.roleInteractor, h.roleIDPInteractor)
	de, err := rFactory.Create(ctx, role.Name, role.Description, scopes, domain.SystemUUID)
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

func mapUserScopeToModel(scope enums.UserScope) model.UserPermision {
	switch scope {
	case enums.ROLES_READ:
		return model.UserPermisionRoleRead
	case enums.ROLES_WRITE:
		return model.UserPermisionRoleWrite
	case enums.USERS_READ:
		return model.UserPermisionUserRead
	case enums.USERS_WRITE:
		return model.UserPermisionUserWrite
	}
	return ""
}

func mapDomainRoleToModel(role *domain.Role) *model.Role {
	scopes := make([]model.UserPermision, 0, len(role.Scopes))
	for _, s := range role.Scopes {
		scopes = append(scopes, mapUserScopeToModel(s))
	}
	return &model.Role{
		Name:        role.Name,
		Description: role.Description,
		RolePk:      role.UUID,
		Permissions: scopes,
	}
}
