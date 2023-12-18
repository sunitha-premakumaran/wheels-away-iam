package zitadel_idp

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/sunitha/wheels-away-iam/internal/core/domain"
	"github.com/sunitha/wheels-away-iam/pkg/zitadel"
	"github.com/zitadel/zitadel-go/pkg/client/zitadel/management"
)

type ZitadelUserInteractor struct {
	*zitadel.ZitadelClient
	logger *zerolog.Logger
}

func NewZitadelUserInteractor(client *zitadel.ZitadelClient, logger *zerolog.Logger) *ZitadelUserInteractor {
	return &ZitadelUserInteractor{
		ZitadelClient: client,
		logger:        logger,
	}
}

func (c *ZitadelUserInteractor) CreateIDPUser(ctx context.Context, user *domain.User) (string, error) {
	userZ := mapDomainUserToZitadel(user)
	response, err := c.Client.ImportHumanUser(ctx, userZ)
	if err != nil {
		c.logger.Error().Msgf("error while creating user in zitadel: %s", err.Error())
		return "", fmt.Errorf("error while creating user in zitadel: %w", err)
	}
	return response.UserId, nil
}

func (c *ZitadelUserInteractor) CreateUserGrant(ctx context.Context, userID string, roles []string) error {
	grantZ := mapToUserGrant(c.ProjectID, userID, roles)
	_, err := c.Client.AddUserGrant(ctx, grantZ)
	if err != nil {
		c.logger.Error().Msgf("error while creating user in zitadel: %s", err.Error())
		return fmt.Errorf("error while creating user in zitadel: %w", err)
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
