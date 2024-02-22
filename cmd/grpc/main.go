package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/sunitha/wheels-away-iam/config"
	"github.com/sunitha/wheels-away-iam/internal/core/services"
	grpc_handler "github.com/sunitha/wheels-away-iam/internal/handlers/grpc"
	"github.com/sunitha/wheels-away-iam/internal/infrastructure/repository"
	"github.com/sunitha/wheels-away-iam/pkg/gorm"
	"github.com/sunitha/wheels-away-iam/pkg/logger"
	pb "github.com/sunitha/wheels-away-iam/protos"
	"google.golang.org/grpc"
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

	userRepo := repository.NewUserRepository(gc.DB)
	userInteractor := services.NewUserInteractor(userRepo)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GRPCPort))
	if err != nil {
		logger.Panic().AnErr("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	userProcessor := grpc_handler.NewUserGrpcHandler(logger, userInteractor)
	pb.RegisterUserProcessorServer(server, userProcessor)
	go func() {
		logger.Printf("server listening at %v", lis.Addr())
		if err := server.Serve(lis); err != nil {
			logger.Fatal().AnErr("failed to serve: %v", err)
		}
	}()

	gracefulShutdown(server)
}

func gracefulShutdown(srv *grpc.Server) {
	// the duration for which the server gracefully wait for existing connections to finish
	var wait = time.Second * 15

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	_, cancel := context.WithTimeout(context.Background(), wait)
	// Note, defers are called LIFO order.
	defer os.Exit(0)
	defer fmt.Println("shutting down api")
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.

	// close grpc server
	srv.GracefulStop()
}
