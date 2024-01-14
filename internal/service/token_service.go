package service

import (
	"context"
	"github.com/ccesarfp/shrine/internal/model"
	"github.com/ccesarfp/shrine/internal/protobuf"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

var (
	jwtSecretKey = "JWT_SECRET_KEY"
)

type Server struct {
	protobuf.UnimplementedTokenServer
}

func (s *Server) CreateToken(ctx context.Context, in *protobuf.UserRequest) (*protobuf.TokenResponse, error) {
	u := model.NewUser(in.Id, in.AppOrigin, in.AccessLevel, in.HoursToExpire)

	claims := jwt.MapClaims{
		"id":          u.Id(),
		"appOrigin":   u.AppOrigin(),
		"accessLevel": u.AccessLevel(),
		"exp":         time.Now().Add(time.Hour * time.Duration(u.HoursToExpire())).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv(jwtSecretKey)))
	if err != nil {
		log.Panicln("Error trying to generate JWT token, err=", err.Error())
	}

	return &protobuf.TokenResponse{
		Token: tokenString,
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
