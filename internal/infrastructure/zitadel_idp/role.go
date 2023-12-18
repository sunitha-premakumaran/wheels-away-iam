package zitadel_idp

import (
	"context"
	"fmt"
	"strings"

	"github.com/rs/zerolog"
	"github.com/sunitha/wheels-away-iam/internal/core/domain"
	"github.com/sunitha/wheels-away-iam/pkg/zitadel"
	"github.com/zitadel/zitadel-go/pkg/client/zitadel/management"
)

type ZitadelRoleInteractor struct {
	*zitadel.ZitadelClient
	logger *zerolog.Logger
}

func NewZitadelRoleInteractor(client *zitadel.ZitadelClient, logger *zerolog.Logger) *ZitadelRoleInteractor {
	return &ZitadelRoleInteractor{
		ZitadelClient: client,
		logger:        logger,
	}
}

func (c *ZitadelUserInteractor) CreateIDPRole(ctx context.Context, role *domain.Role) error {
	roleZ := mapDomainRoleToZitadel(c.ProjectID, role)
	_, err := c.Client.AddProjectRole(ctx, roleZ)
	if err != nil {
		c.logger.Error().Msgf("error while creating role in zitadel: %s", err.Error())
		return fmt.Errorf("error while creating role in zitadel: %w", err)
	}
	return nil
}

func mapDomainRoleToZitadel(projectID string, role *domain.Role) *management.AddProjectRoleRequest {
	return &management.AddProjectRoleRequest{
		ProjectId:   projectID,
		RoleKey:     strings.ToLower(role.Name),
		DisplayName: role.Name,
	}
}
