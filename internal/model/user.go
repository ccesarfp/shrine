package model

// User Model
// id - user id
// name - user name
// appOrigin - name of the application responsible for the user
// accessLevel - user level access
// token - user token
// hoursToExpire - token expiration time
// **
type User struct {
	id            int64
	name          string
	appOrigin     string
	accessLevel   int32
	token         string
	hoursToExpire int32
}

func NewUser(id int64, name string, appOrigin string, accessLevel int32, hoursToExpire int32) User {
	return User{
		id:            id,
		name:          name,
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

func (u *User) Name() string {
	return u.name
}

func (u *User) AccessLevel() int32 {
	return u.accessLevel
}

func (u *User) HoursToExpire() int32 {
	return u.hoursToExpire
}
