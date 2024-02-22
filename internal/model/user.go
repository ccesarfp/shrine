package model

import (
	"github.com/go-playground/validator/v10"
)

// User Model
// Token - user token
// HoursToExpire - token expiration time
// **
type User struct {
	IpAddress     string `validate:"required"`
	Token         string `validate:"omitempty,uuid5"`
	HoursToExpire int32  `validate:"required,min=1"`
}

func NewUser(ipAddress string, hoursToExpire int32) (*User, error) {
	// Setting default value for hoursToExpire
	if hoursToExpire == 0 {
		hoursToExpire = 1
	}

	u := User{
		IpAddress:     ipAddress,
		HoursToExpire: hoursToExpire,
	}

	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
