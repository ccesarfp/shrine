package service

import (
	"context"
	"github.com/ccesarfp/shrine/internal/config"
	"github.com/ccesarfp/shrine/internal/errors"
	"github.com/ccesarfp/shrine/internal/model"
	"github.com/ccesarfp/shrine/internal/protobuf"
	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"os"
	"time"
)

var (
	otSecretKey = "OT_SECRET_KEY"
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
	uuidValue, err := uuid.FromString(os.Getenv(otSecretKey))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Getting actual datetime
	currentTime := time.Now()

	// Getting IP Address from Request
	p, _ := peer.FromContext(ctx)
	request := p.Addr.String()

	opaqueToken, err := model.NewOpaqueToken(uuid.NewV5(uuidValue, request+currentTime.String()).String())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	u, err := model.NewUser(in.Id, in.Name, in.AppOrigin, in.AccessLevel, in.HoursToExpire)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	exp := time.Now().Add(time.Hour * time.Duration(u.HoursToExpire))

	claims := jwt.MapClaims{
		"id":          u.Id,
		"name":        u.Name,
		"appOrigin":   u.AppOrigin,
		"accessLevel": u.AccessLevel,
		"exp":         exp.Unix(),
	}

	// Creating Jwt
	t := model.Jwt{}
	token, err := t.CreateJwt(claims, os.Getenv(otSecretKey))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Creating Redis Client instance
	client, err := config.NewRedisClient()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Writing token to db
	err = client.Set(ctx, opaqueToken.Token, token, exp.Sub(time.Now())).Err()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protobuf.TokenResponse{
		Token: opaqueToken.Token,
	}, nil
}

// GetClaimsByKey Retrieve data from JWT using Jwt ID
// params:
//   - TokenRequestWithId - token id
//
// result:
//   - UserResponseWithToken - user data and token
//
// **
func (s *Server) GetClaimsByKey(ctx context.Context, in *protobuf.TokenRequestWithId) (*protobuf.UserResponseWithToken, error) {
	op, err := model.NewOpaqueToken(in.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Creating Redis Client instance
	client, err := config.NewRedisClient()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Searching token in db
	jwtString, err := client.Get(ctx, op.Token).Result()
	if err != nil {
		// If the token does not exist in the db, returns Not Found
		if err.Error() == "redis: nil" {
			return nil, status.Error(codes.NotFound, "token does not exist")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	op.SetJwt(jwtString)

	// If token is valid, return claims
	if op.Jwt != "" {
		return &protobuf.UserResponseWithToken{
			Token: op.Jwt,
		}, nil
	}

	return nil, status.Error(codes.Unknown, "Error")
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
	t, err := model.NewJwt(in.Token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Getting claims
	token, claims, err := t.GetClaims(otSecretKey)
	if err != nil {
		// If token is not valid, return Unauthenticated
		if token.Valid == false {
			expiredToken := errors.ExpiredToken{}
			return nil, status.Error(codes.Unauthenticated, expiredToken.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	// If token is valid, return claims
	if token.Valid {
		return &protobuf.UserResponse{
			Id:          int64(claims["id"].(float64)),
			Name:        claims["name"].(string),
			AppOrigin:   claims["appOrigin"].(string),
			AccessLevel: int32(claims["accessLevel"].(float64)),
		}, nil
	}

	return nil, status.Error(codes.Unknown, "Error")
}

// CheckTokenValidity Verify token validity
// params:
//   - TokenRequest - user token
//
// result:
//   - TokenStatus - token status
//
// **
func (s *Server) CheckTokenValidity(ctx context.Context, in *protobuf.TokenRequest) (*protobuf.TokenStatus, error) {
	t, err := model.NewJwt(in.Token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	isValid, err := t.CheckValidity(otSecretKey)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return &protobuf.TokenStatus{Status: isValid}, nil
}
