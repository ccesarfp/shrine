package service

import (
	"context"
	"github.com/ccesarfp/shrine/internal/config/redis"
	"github.com/ccesarfp/shrine/internal/errors"
	jwt2 "github.com/ccesarfp/shrine/internal/model/jwt"
	"github.com/ccesarfp/shrine/internal/model/opaque_token"
	"github.com/ccesarfp/shrine/internal/model/user"
	"github.com/ccesarfp/shrine/internal/protobuf"
	"github.com/ccesarfp/shrine/pkg/util"
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

// CreateToken Create user opaque token
// params:
//   - UserRequest - user data
//
// result:
//   - UserResponse - opaque token created
//
// **
func (s *Server) CreateToken(ctx context.Context, in *protobuf.UserRequest) (*protobuf.UserResponse, error) {
	// Getting Environment Secret
	secret := os.Getenv(otSecretKey)

	// Creating UUID
	uuidValue, err := uuid.FromString(secret)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Getting actual datetime
	currentTime := time.Now()

	// Getting IP Address from Request
	p, _ := peer.FromContext(ctx)
	ipAddress := p.Addr.String()

	// Creating User
	u, err := user.New(ipAddress, in.HoursToExpire)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Creating Opaque Token
	opaqueToken, err := opaque_token.New(uuid.NewV5(uuidValue, u.IpAddress+currentTime.String()).String())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Creating Claims
	exp, err := util.CreateUnixExpirationTime(u.HoursToExpire)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	claims := jwt.MapClaims{
		"ipAddress": u.IpAddress,
		"exp":       exp.Unix(),
	}
	t := jwt2.Jwt{}
	jwtString, err := t.CreateJwt(claims, secret)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Creating Redis Client instance
	client, err := redis.NewRedisClient()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Writing token to db
	err = client.Set(ctx, opaqueToken.Token, jwtString, exp.Sub(time.Now())).Err()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protobuf.UserResponse{
		Token: opaqueToken.Token,
	}, nil
}

// UpdateToken Update user jwt
// params:
//   - UserUpdateRequest - user opaque token and jwt
//
// result:
//   - UserResponse - user opaque token
//
// **
func (s *Server) UpdateToken(ctx context.Context, in *protobuf.UserUpdateRequest) (*protobuf.UserResponse, error) {
	// Creating Opaque Token
	op, err := opaque_token.NewWithJwt(in.Token, in.Jwt)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Creating Redis Client instance
	client, err := redis.NewRedisClient()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Preparing token validity
	exp, err := util.CreateUnixExpirationTime(in.HoursToExpire)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Updating Jwt
	clientResponse := client.SetXX(ctx, op.Token, op.Jwt, exp.Sub(time.Now()))
	err = clientResponse.Err()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// If the token does not exist, return error
	if clientResponse.Val() == false {
		err = &errors.ExpiredToken{}
		return nil, status.Error(codes.NotFound, err.Error())
	}

	// Returning Token
	return &protobuf.UserResponse{
		Token: op.Token,
	}, nil
}

// GetJwt Returns JWT using Opaque Token
// params:
//   - TokenRequest - opaque token
//
// result:
//   - TokenResponse - user jwt
//
// **
func (s *Server) GetJwt(ctx context.Context, in *protobuf.TokenRequest) (*protobuf.TokenResponse, error) {
	// Creating Opaque Token
	ot, err := opaque_token.New(in.Token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Creating Redis Client instance
	client, err := redis.NewRedisClient()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Searching token in db
	jwtString, err := client.Get(ctx, ot.Token).Result()
	if err != nil {
		// If the token does not exist in the db, returns Not Found
		if err.Error() == "redis: nil" {
			return nil, status.Error(codes.NotFound, "token does not exist")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	ot.SetJwt(jwtString)

	// If token is valid, return claims
	if ot.Jwt != "" {
		return &protobuf.TokenResponse{
			Jwt: ot.Jwt,
		}, nil
	}

	return nil, status.Error(codes.Unknown, "Error")
}

// CheckTokenValidity verify token validity
// params:
//   - TokenRequest - user token
//
// result:
//   - TokenStatus - token status
//
// **
func (s *Server) CheckTokenValidity(ctx context.Context, in *protobuf.TokenRequest) (*protobuf.TokenStatus, error) {
	// Creating Opaque Token
	ot, err := opaque_token.New(in.Token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Creating Redis Client instance
	client, err := redis.NewRedisClient()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Verifying token in db
	statusToken := client.Exists(ctx, ot.Token)

	// Preparing return status
	isValid := false
	if statusToken.Val() == 1 {
		isValid = true
	}

	return &protobuf.TokenStatus{Status: isValid}, nil
}
