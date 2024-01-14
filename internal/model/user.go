package model

type User struct {
	id            int64
	appOrigin     string
	accessLevel   int32
	token         string
	hoursToExpire int32
}

func NewUser(id int64, appOrigin string, accessLevel int32, hoursToExpire int32) User {
	return User{
		id:            id,
		appOrigin:     appOrigin,
		accessLevel:   accessLevel,
		hoursToExpire: hoursToExpire,
	}
}

func (u *User) Id() int64 {
	return u.id
}

func (u *User) AppOrigin() string {
	return u.appOrigin
}

func (u *User) AccessLevel() int32 {
	return u.accessLevel
}

func (u *User) HoursToExpire() int32 {
	return u.hoursToExpire
}
