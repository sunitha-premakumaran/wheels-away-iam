package grpc

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	pb "github.com/sunitha/wheels-away-iam/protos"
)

type UserGrpcHandler struct {
	logger         *zerolog.Logger
	userInteractor UserInteractor

	pb.UnimplementedUserProcessorServer
}

func NewUserGrpcHandler(logger *zerolog.Logger, userInteractor UserInteractor) *UserGrpcHandler {
	return &UserGrpcHandler{
		logger:         logger,
		userInteractor: userInteractor,
	}
}

func (s *UserGrpcHandler) GetUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	user, err := s.userInteractor.GetUser(ctx, in.GetId())
	if err != nil {
		return nil, fmt.Errorf("error while getting ")
	}
	return &pb.UserResponse{
		UserId:    user.User.UUID,
		UserName:  fmt.Sprintf("%s, %s", user.User.LastName, user.User.FirstName),
		UserEmail: user.User.Email,
	}, nil
}
