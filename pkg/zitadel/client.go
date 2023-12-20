package zitadel

import (
	"github.com/rs/zerolog"
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
	jwt := []byte(config.JWTToken)
	client, err := management.NewClient(config.Scopes,
		zitadel.WithCustomURL(config.Issuer, config.URL),
		zitadel.WithJWTProfileTokenSource(
			middleware.JWTProfileFromFileData(jwt),
		))
	if err != nil {
		logger.Fatal().AnErr("could not create client", err)
	}
	logger.Info().Msgf("connection to zitadel successful")
	return &ZitadelClient{
		Client:    client,
		ProjectID: config.ProjectID,
		Logger:    logger,
	}
}
