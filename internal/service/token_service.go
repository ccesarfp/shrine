package service

import (
	"context"
	"fmt"
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

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(t.Token(), claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv(jwtSecretKey)), nil
	})

	if err != nil {
		fmt.Println("Error parsing token, err=:", err)
	}

	if token.Valid {
		return &protobuf.UserResponse{
			Id:          int64(claims["id"].(float64)),
			Name:        claims["name"].(string),
			AppOrigin:   claims["appOrigin"].(string),
			AccessLevel: int32(claims["accessLevel"].(float64)),
		}, nil
	}

	return &protobuf.UserResponse{}, nil
}
