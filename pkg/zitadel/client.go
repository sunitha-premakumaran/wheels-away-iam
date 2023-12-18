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
}

func NewZitadelClient(config Config, logger *zerolog.Logger) *ZitadelClient {
	client, err := management.NewClient(config.Scopes,
		zitadel.WithCustomURL(config.Issuer, config.URL),
		zitadel.WithJWTProfileTokenSource(
			middleware.JWTProfileFromFileData(config.JWTToken),
		))
	if err != nil {
		logger.Fatal().AnErr("could not create client", err)
	}
	defer func() {
		err := client.Connection.Close()
		if err != nil {
			logger.Error().Msgf("could not close grpc connection: %s", err.Error())
		}
	}()
	return &ZitadelClient{Client: client, ProjectID: config.ProjectID}
}
