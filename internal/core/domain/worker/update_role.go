package worker

import (
	"context"
	"fmt"

	"github.com/sunitha/wheels-away-iam/internal/core/domain"
	"github.com/sunitha/wheels-away-iam/internal/core/enums"
)

type UpdateRoleFactory struct {
	roleUserMappingInteractor RoleUserMappingInteractor
	roleInteractor            RoleInteractor
}

func NewUpdateRoleFactory(
	roleInteractor RoleInteractor,
	roleUserMappingInteractor RoleUserMappingInteractor,
) *UpdateRoleFactory {
	return &UpdateRoleFactory{
		roleInteractor:            roleInteractor,
		roleUserMappingInteractor: roleUserMappingInteractor,
	}
}

func (w UpdateRoleFactory) Do(ctx context.Context, rolePK, name string, description *string, scopes []enums.UserScope) (domain.HttpErrorCode, error) {
	role, err := w.roleInteractor.GetRole(ctx, rolePK)
	if err != nil {
		return domain.InternalServerError, err
	}
	if role == nil {
		return domain.BadRequestError, fmt.Errorf("%s role not found", rolePK)
	}
	role.Name = name
	role.Description = description
	role.Scopes = scopes
	err = w.roleInteractor.SaveRole(ctx, role)
	if err != nil {
		return domain.InternalServerError, err
	}
	return "", nil
}
