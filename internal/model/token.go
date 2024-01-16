package model

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
)

// Token Model
// Id - token Id
// token - user token
// **
type Token struct {
	Id    string
	Token string `validate:"omitempty,jwt"`
}

func NewToken(token string) Token {
	return Token{
		Token: token,
	}
}

func NewTokenWithId(id string) Token {
	return Token{
		Id: id,
	}
}

func (t *Token) SetToken(token string) {
	t.Token = token
}

// CreateToken create a token
// params:
//   - claims jwt.MapClaims - claims
//   - jwtSecretKey string - secret that will be used to create the token
//
// returns:
//   - tokenString string - token created
//   - err error - error message
//
// **
func (t *Token) CreateToken(claims jwt.MapClaims, jwtSecretKey string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv(jwtSecretKey)))
	if err != nil {
		log.Println("Error trying to generate JWT token, err=", err.Error())
		return "nil", err
	}

	return tokenString, nil
}

// GetClaims accesses a token and returns its claims and validity
// params:
//   - jwtSecretKey string - secret used to create the token
//
// returns:
//   - token *jwt.Token - token and its validity
//   - claims jwt.MapClaims - token claims
//   - err error - error message
//
// **
func (t *Token) GetClaims(jwtSecretKey string) (*jwt.Token, jwt.MapClaims, error) {

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(t.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv(jwtSecretKey)), nil
	})

	if err != nil {
		return token, nil, err
	}

	return token, claims, nil

}

func (t *Token) CheckValidity(jwtSecretKey string) (bool, error) {

	isValid, _, err := t.GetClaims(jwtSecretKey)
	if err != nil {
		return false, err
	}

	return isValid.Valid, nil
}
