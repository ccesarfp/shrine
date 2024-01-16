package model

import (
	"github.com/go-playground/validator/v10"
)

// User Model
// Id - user Id
// Name - Name of user
// AppOrigin - Name of the application responsible for the user
// AccessLevel - user level access
// token - user token
// hoursToExpire - token expiration time
// **
type User struct {
	Id            int64  `validate:"required"`
	Name          string `validate:"required"`
	AppOrigin     string `validate:"required"`
	AccessLevel   int32  `validate:"required,min=1"`
	Token         string `validate:"omitempty,jwt"`
	HoursToExpire int32  `validate:"required,min=1"`
}

func NewUser(id int64, name string, appOrigin string, accessLevel int32, hoursToExpire int32) (*User, error) {
	u := User{
		Id:            id,
		Name:          name,
		AppOrigin:     appOrigin,
		AccessLevel:   accessLevel,
		HoursToExpire: hoursToExpire,
	}

	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
