package zitadel_idp

import (
	"context"
	"fmt"
	"strings"

	"github.com/sunitha/wheels-away-iam/internal/core/domain"
	"github.com/sunitha/wheels-away-iam/pkg/zitadel"
	"github.com/zitadel/zitadel-go/pkg/client/zitadel/management"
)

type ZitadelRoleInteractor struct {
	connection *zitadel.ZitadelClient
}

func NewZitadelRoleInteractor(client *zitadel.ZitadelClient) *ZitadelRoleInteractor {
	return &ZitadelRoleInteractor{
		connection: client,
	}
}

func (c *ZitadelRoleInteractor) SaveIDPRole(ctx context.Context, role *domain.Role) (string, error) {
	roleZ := mapDomainRoleToZitadel(c.connection.ProjectID, role)
	_, err := c.connection.Client.AddProjectRole(ctx, roleZ)
	if err != nil {
		c.connection.Logger.Error().Msgf("error while creating role in zitadel: %s", err.Error())
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
