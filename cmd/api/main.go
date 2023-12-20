package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog"
	"github.com/sunitha/wheels-away-iam/config"
	"github.com/sunitha/wheels-away-iam/graph"
	"github.com/sunitha/wheels-away-iam/graph/generated"
	"github.com/sunitha/wheels-away-iam/internal/core/services"
	"github.com/sunitha/wheels-away-iam/internal/handlers/api"
	"github.com/sunitha/wheels-away-iam/internal/infrastructure/repository"
	"github.com/sunitha/wheels-away-iam/internal/infrastructure/zitadel_idp"
	"github.com/sunitha/wheels-away-iam/pkg/gorm"
	"github.com/sunitha/wheels-away-iam/pkg/health"
	"github.com/sunitha/wheels-away-iam/pkg/logger"
	"github.com/sunitha/wheels-away-iam/pkg/middlewares"
	"github.com/sunitha/wheels-away-iam/pkg/zitadel"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}
	configName := fmt.Sprintf("config.%s", env)
	config := config.Init(configName)

	logger := logger.NewLogger(config)
	gc := gorm.NewDBClient(&config.Database, logger)

	zClient := zitadel.NewZitadelClient(config.ZitadelConfig, logger)

	roleRepo := repository.NewRoleRepository(gc.DB)
	roleInteractor := services.NewRoleInteractor(roleRepo)

	roleUserMapRepo := repository.NewRoleUserMappingRepository(gc.DB)
	roleUserMapInteractor := services.NewRoleUserMappingInteractor(roleUserMapRepo)

	userIDPInteractor := zitadel_idp.NewZitadelUserInteractor(zClient, logger)

	userRepo := repository.NewUserRepository(gc.DB)
	userInteractor := services.NewUserInteractor(userRepo)
	userProcessor := api.NewUserProcessor(userInteractor, userIDPInteractor, roleUserMapInteractor, roleInteractor)

	mux := http.NewServeMux()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		UserProcessor: userProcessor,
	}}))

	mux.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/healthz", http.HandlerFunc(health.HealthCheck))
	mux.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		cm := middlewares.NewMiddleware(logger)
		middlewares.Chain(srv, cm).ServeHTTP(w, r)
	})

	server := &http.Server{
		Addr:         fmt.Sprintf(":%v", config.APIPort),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      mux,
	}

	go func() {
		logger.Info().Msgf("connect to http://localhost:%d/playground for GraphQL playground", config.APIPort)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Panic().Msgf("Failed to start %v", err)
		}
	}()

	gracefulShutdown(server, logger)
}

func gracefulShutdown(server *http.Server, logger *zerolog.Logger) {
	// the duration for which the server gracefully wait for existing connections to finish
	var wait = time.Second * 15

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	// Note, defers are called LIFO order.
	defer os.Exit(0)
	defer logger.Print("shutting down api")
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.

	// close server
	err := server.Shutdown(ctx)
	if err != nil {
		logger.Fatal().Msgf("error while shutting down api: %v", err)
	}
}
