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
		Status: "ok",
		Token:  "aaaaa",
	}, nil
}
