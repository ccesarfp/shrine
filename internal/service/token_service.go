package service

import (
	"context"
	"github.com/ccesarfp/shrine/internal/protobuf"
)

type Server struct {
	protobuf.UnimplementedTokenServer
}

func (s *Server) CreateToken(ctx context.Context, in *protobuf.UserRequest) (*protobuf.TokenResponse, error) {
	return &protobuf.TokenResponse{
		Token: "aaaaa",
	}, nil
}

func (s *Server) GetClaimsByKey(ctx context.Context, in *protobuf.TokenRequestWithId) (*protobuf.UserResponseWithToken, error) {
	return &protobuf.UserResponseWithToken{
		Token: "aaaaa",
	}, nil
}

func (s *Server) GetClaimsByToken(ctx context.Context, in *protobuf.TokenRequest) (*protobuf.UserResponse, error) {
	return &protobuf.UserResponse{
		Id: 1,
	}, nil
}
