package zitadel

import (
	"github.com/rs/zerolog"
	"github.com/zitadel/oidc/pkg/oidc"
	"github.com/zitadel/zitadel-go/pkg/client/management"
	"github.com/zitadel/zitadel-go/pkg/client/middleware"
	"github.com/zitadel/zitadel-go/pkg/client/zitadel"
)

type ZitadelClient struct {
	Client    *management.Client
	ProjectID string
	Logger    *zerolog.Logger
}

func NewZitadelClient(config Config, logger *zerolog.Logger) *ZitadelClient {
	client, err := management.NewClient([]string{oidc.ScopeOpenID, zitadel.ScopeProjectID(config.ProjectID), "urn:zitadel:iam:org:project:id:zitadel:aud"},
		zitadel.WithInsecure(),
		zitadel.WithCustomURL(config.Issuer, config.URL),
		zitadel.WithJWTProfileTokenSource(
			middleware.JWTProfileFromPath("jwt-key.json"),
		),
		zitadel.WithOrgID("245155221376925699"),
	)
	if err != nil {
		logger.Panic().Msgf("could not create client: %s", err.Error())
	}
	logger.Info().Msgf("zitadel client connection successful")
	return &ZitadelClient{
		Client:    client,
		ProjectID: config.ProjectID,
		Logger:    logger,
	}
}
