package zitadel_idp

import (
	"context"
	"fmt"

	"github.com/sunitha/wheels-away-iam/internal/core/domain"
	"github.com/sunitha/wheels-away-iam/pkg/zitadel"
	"github.com/zitadel/zitadel-go/pkg/client/zitadel/management"
)

type ZitadelUserInteractor struct {
	*zitadel.ZitadelClient
}

func NewZitadelUserInteractor(client *zitadel.ZitadelClient) *ZitadelUserInteractor {
	return &ZitadelUserInteractor{
		ZitadelClient: client,
	}
}

func (c *ZitadelUserInteractor) CreateIDPUser(ctx context.Context, user *domain.User) (string, error) {
	userZ := mapDomainUserToZitadel(user)
	response, err := c.Client.ImportHumanUser(ctx, userZ)
	if err != nil {
		c.Logger.Error().Msgf("error while creating user in zitadel: %s", err.Error())
		return "", fmt.Errorf("error while creating user in zitadel: %w", err)
	}
	return response.UserId, nil
}

func (c *ZitadelUserInteractor) UpdateIDPUser(ctx context.Context, user *domain.User) error {
	userZ := mapDomainUserToZitadelProfile(user)
	_, err := c.Client.UpdateHumanProfile(ctx, userZ)
	if err != nil {
		c.Logger.Error().Msgf("error while updating user_profile in zitadel: %s", err.Error())
		return fmt.Errorf("error while updating user_profile in zitadel: %w", err)
	}
	_, err = c.Client.UpdateHumanEmail(ctx, &management.UpdateHumanEmailRequest{
		Email: user.Email,
	})
	if err != nil {
		c.Logger.Error().Msgf("error while updating user_email in zitadel: %s", err.Error())
		return fmt.Errorf("error while updating user_email in zitadel: %w", err)
	}
	_, err = c.Client.UpdateHumanPhone(ctx, &management.UpdateHumanPhoneRequest{
		Phone: user.Phone,
	})
	if err != nil {
		c.Logger.Error().Msgf("error while updating user_phone in zitadel: %s", err.Error())
		return fmt.Errorf("error while updating user_phone in zitadel: %w", err)
	}
	return nil
}

func (c *ZitadelUserInteractor) CreateUserGrant(ctx context.Context, userID string, roles []string) (string, error) {
	grantZ := mapToUserGrant(c.ProjectID, userID, roles)
	res, err := c.Client.AddUserGrant(ctx, grantZ)
	if err != nil {
		c.Logger.Error().Msgf("error while creating user_grant in zitadel: %s", err.Error())
		return "", fmt.Errorf("error while creating user_grant in zitadel: %w", err)
	}
	return res.UserGrantId, nil
}

func (c *ZitadelUserInteractor) DeleteUserGrant(ctx context.Context, userID string, grantID string) error {
	_, err := c.Client.RemoveUserGrant(ctx, &management.RemoveUserGrantRequest{UserId: userID, GrantId: grantID})
	if err != nil {
		c.Logger.Error().Msgf("error while creating user_grant in zitadel: %s", err.Error())
		return fmt.Errorf("error while creating user_grant in zitadel: %w", err)
	}
	return nil
}

func mapToUserGrant(projectID, userID string, roles []string) *management.AddUserGrantRequest {
	return &management.AddUserGrantRequest{
		UserId:    userID,
		RoleKeys:  roles,
		ProjectId: projectID,
	}
}

func mapDomainUserToZitadel(user *domain.User) *management.ImportHumanUserRequest {
	return &management.ImportHumanUserRequest{
		UserName: user.Email,
		Profile: &management.ImportHumanUserRequest_Profile{
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
		Email: &management.ImportHumanUserRequest_Email{
			Email:           user.Email,
			IsEmailVerified: false,
		},
		Phone: &management.ImportHumanUserRequest_Phone{
			Phone:           user.Phone,
			IsPhoneVerified: false,
		},
		Password: user.Password,
	}
}

func mapDomainUserToZitadelProfile(user *domain.User) *management.UpdateHumanProfileRequest {
	return &management.UpdateHumanProfileRequest{
		UserId:    user.AuthID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}
