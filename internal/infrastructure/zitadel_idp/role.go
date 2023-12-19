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

func (c *ZitadelUserInteractor) SaveIDPRole(ctx context.Context, role *domain.Role) (string, error) {
	roleZ := mapDomainRoleToZitadel(c.ProjectID, role)
	_, err := c.Client.AddProjectRole(ctx, roleZ)
	if err != nil {
		c.logger.Error().Msgf("error while creating role in zitadel: %s", err.Error())
		return "", fmt.Errorf("error while creating role in zitadel: %w", err)
	}
	return roleZ.RoleKey, nil
}

func mapDomainRoleToZitadel(projectID string, role *domain.Role) *management.AddProjectRoleRequest {
	rn := strings.ToLower(strings.Join(strings.Split(role.Name, " "), "-"))
	return &management.AddProjectRoleRequest{
		ProjectId:   projectID,
		RoleKey:     rn,
		DisplayName: role.Name,
	}
}
