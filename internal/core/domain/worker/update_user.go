package worker

import (
	"context"

	"github.com/sunitha/wheels-away-iam/internal/core/domain"
	"github.com/sunitha/wheels-away-iam/internal/core/enums"
)

type UpdateUserWorker struct {
	userInteractor            UserInteractor
	userIDPInteractor         UserIDPInteractor
	roleUserMappingInteractor RoleUserMappingInteractor
	roleInteractor            RoleInteractor
}

func NewUpdateUserWorker(
	userInteractor UserInteractor,
	userIDPInteractor UserIDPInteractor,
	roleInteractor RoleInteractor,
	roleUserMappingInteractor RoleUserMappingInteractor,
) *UpdateUserWorker {
	return &UpdateUserWorker{
		userInteractor:            userInteractor,
		userIDPInteractor:         userIDPInteractor,
		roleInteractor:            roleInteractor,
		roleUserMappingInteractor: roleUserMappingInteractor,
	}
}

func (w *UpdateUserWorker) Do(ctx context.Context, userID, firstName, lastName, email, phone string, profileImageUrl, description *string,
	status enums.UserStatus, userRoleIDs []string, updatedBy string) (domain.HttpErrorCode, error) {
	user, err := w.userInteractor.GetUser(ctx, userID)
	if err != nil {
		return domain.InternalServerError, err
	}
	if user == nil {
		return domain.BadRequestError, err
	}
	user.User.FirstName = firstName
	user.User.LastName = lastName
	user.User.Email = email
	user.User.Phone = phone
	user.User.ProfileImage = profileImageUrl
	user.User.Description = description
	user.User.Status = status
	err = w.userInteractor.SaveUser(ctx, user.User)
	if err != nil {
		return domain.InternalServerError, err
	}
	err = w.userIDPInteractor.UpdateIDPUser(ctx, user.User)
	if err != nil {
		return domain.InternalServerError, err
	}

	return "", nil
}
