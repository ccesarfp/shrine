package service

import (
	"context"
	"github.com/ccesarfp/shrine/internal/model"
	"github.com/ccesarfp/shrine/internal/protobuf"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

var (
	jwtSecretKey = "JWT_SECRET_KEY"
)

type Server struct {
	protobuf.UnimplementedTokenServer
}

// CreateToken Create user token
// params:
//   - UserRequest - user data
//
// result:
//   - TokenResponse - token created
//
// **
func (s *Server) CreateToken(ctx context.Context, in *protobuf.UserRequest) (*protobuf.TokenResponse, error) {
	u := model.NewUser(in.Id, in.Name, in.AppOrigin, in.AccessLevel, in.HoursToExpire)

	claims := jwt.MapClaims{
		"id":          u.Id(),
		"name":        u.Name(),
		"appOrigin":   u.AppOrigin(),
		"accessLevel": u.AccessLevel(),
		"exp":         time.Now().Add(time.Hour * time.Duration(u.HoursToExpire())).Unix(),
	}

	token, err := model.Token{}.CreateToken(claims, jwtSecretKey)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protobuf.TokenResponse{
		Token: token,
	}, nil
}

func (s *Server) GetClaimsByKey(ctx context.Context, in *protobuf.TokenRequestWithId) (*protobuf.UserResponseWithToken, error) {
	return &protobuf.UserResponseWithToken{
		Token: "aaaaa",
	}, nil
}

// GetClaimsByToken Retrieve data from JWT
// params:
//   - TokenRequest - user token
//
// result:
//   - UserResponse - user data
//
// **
func (s *Server) GetClaimsByToken(ctx context.Context, in *protobuf.TokenRequest) (*protobuf.UserResponse, error) {
	t := model.NewToken(in.Token)

	token, claims, err := t.GetClaims(jwtSecretKey)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if token.Valid {
		return &protobuf.UserResponse{
			Id:          int64(claims["id"].(float64)),
			Name:        claims["name"].(string),
			AppOrigin:   claims["appOrigin"].(string),
			AccessLevel: int32(claims["accessLevel"].(float64)),
		}, nil
	}

	return nil, status.Error(codes.Unauthenticated, "the token has expired")
}

func (s *Server) CheckTokenValidity(ctx context.Context, in *protobuf.TokenRequest) (*protobuf.TokenStatus, error) {

	t := model.NewToken(in.Token)

	isValid, err := t.CheckValidity(jwtSecretKey)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if isValid == false {
		return nil, status.Error(codes.Unauthenticated, "the token has expired")
	}

	return &protobuf.TokenStatus{Status: isValid}, nil
}
