package factory

import (
	"context"
	"fmt"

	"github.com/sunitha/wheels-away-iam/internal/core/domain"
	"github.com/sunitha/wheels-away-iam/internal/core/enums"
)

type CreateUserFactory struct {
	userInteractor            UserInteractor
	userIDPInteractor         UserIDPInteractor
	roleUserMappingInteractor RoleUserMappingInteractor
	roleInteractor            RoleInteractor
}

func NewCreateUserFactory(
	userInteractor UserInteractor,
	userIDPInteractor UserIDPInteractor,
	roleInteractor RoleInteractor,
	roleUserMappingInteractor RoleUserMappingInteractor,
) *CreateUserFactory {
	return &CreateUserFactory{
		userInteractor:            userInteractor,
		userIDPInteractor:         userIDPInteractor,
		roleInteractor:            roleInteractor,
		roleUserMappingInteractor: roleUserMappingInteractor,
	}
}

func (f *CreateUserFactory) Create(ctx context.Context, firstName, lastName, email, phone string, profileImageUrl, description *string,
	status enums.UserStatus, userRoleIDs []string, createdBy string) (domain.HttpErrorCode, error) {
	user, err := domain.NewUser(firstName, lastName, email, phone, profileImageUrl, description, status, "", createdBy)
	if err != nil {
		return "", fmt.Errorf("validation for user failed: %w", err)
	}
	var userID string
	userID, err = f.userIDPInteractor.CreateIDPUser(ctx, user)
	if err != nil {
		return "", err
	}
	user.AuthID = userID
	err = f.userInteractor.SaveUser(ctx, user)
	if err != nil {
		return "", err
	}
	userRoles, err := f.roleInteractor.GetRolesByIDs(ctx, userRoleIDs)
	if err != nil {
		return "", err
	}
	if len(userRoles) != len(userRoleIDs) {
		return domain.BadRequestError, fmt.Errorf("Unknown role in the request")
	}
	for _, r := range userRoles {
		roleMap, err := domain.NewRoleUserMapping(r.UUID, user.UUID, createdBy)
		if err != nil {
			return "", err
		}
		err = f.roleUserMappingInteractor.SaveRoleUserMapping(ctx, roleMap)
		if err != nil {
			return "", err
		}
		err = f.userIDPInteractor.CreateUserGrant(ctx, userID, []string{r.AuthKey})
		if err != nil {
			return "", err
		}
	}
	return "", nil
}
