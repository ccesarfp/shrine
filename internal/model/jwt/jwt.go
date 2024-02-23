package jwt

import (
	"errors"
	"github.com/ccesarfp/shrine/pkg/util"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
)

// Jwt Model
// Id - token Id
// token - user token
// **
type Jwt struct {
	Id    string
	Token string `validate:"omitempty,jwt"`
}

var pattern = "^\\d+-[A-Za-z]+$"

func New(token string) (*Jwt, error) {
	t := Jwt{
		Token: token,
	}

	validate := validator.New()
	err := validate.Struct(t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func NewWithId(id string) (*Jwt, error) {
	isValid, err := util.ValidateUsingRegex(pattern, id)
	if err != nil {
		return nil, err
	}
	if isValid == false {
		return nil, errors.New("id not valid")
	}

	t := Jwt{
		Id: id,
	}

	return &t, nil
}

func (t *Jwt) SetJwt(token string) {
	t.Token = token
}

// CreateJwt create a token
// params:
//   - claims jwt.MapClaims - claims
//   - jwtSecretKey string - secret that will be used to create the token
//
// returns:
//   - tokenString string - token created
//   - err error - error message
//
// **
func (t *Jwt) CreateJwt(claims jwt.MapClaims, secret string) (*string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println("Error trying to generate JWT token, err=", err.Error())
		return nil, err
	}

	return &tokenString, nil
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
func (t *Jwt) GetClaims(jwtSecretKey string) (*jwt.Token, jwt.MapClaims, error) {

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(t.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv(jwtSecretKey)), nil
	})

	if err != nil {
		return token, nil, err
	}

	return token, claims, nil

}

func (t *Jwt) CheckValidity(jwtSecretKey string) (bool, error) {

	isValid, _, err := t.GetClaims(jwtSecretKey)
	if err != nil {
		return false, err
	}

	return isValid.Valid, nil
}
