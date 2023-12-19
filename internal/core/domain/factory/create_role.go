package factory

import (
	"context"

	"github.com/sunitha/wheels-away-iam/internal/core/domain"
	"github.com/sunitha/wheels-away-iam/internal/core/enums"
)

type CreateRoleFactory struct {
	roleInteractor    RoleInteractor
	roleIDPInteractor RoleIDPInteractor
}

func NewCreateRoleFactory(
	roleInteractor RoleInteractor,
	roleIDPInteractor RoleIDPInteractor,
) *CreateRoleFactory {
	return &CreateRoleFactory{
		roleInteractor:    roleInteractor,
		roleIDPInteractor: roleIDPInteractor,
	}
}

func (f *CreateRoleFactory) Create(ctx context.Context, name string, description *string, scopes []enums.UserScope, createdBy string) (domain.HttpErrorCode, error) {
	role, err := domain.NewRole(name, description, scopes, createdBy)
	if err != nil {
		return domain.BadRequestError, err
	}
	key, err := f.roleIDPInteractor.SaveIDPRole(ctx, role)
	if err != nil {
		return domain.InternalServerError, err
	}
	role.AuthKey = key
	err = f.roleInteractor.SaveRole(ctx, role)
	if err != nil {
		return domain.InternalServerError, err
	}
	return "", nil
}
