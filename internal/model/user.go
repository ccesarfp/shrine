package model

type User struct {
	simpleModelImpl
	appOrigin     string
	accessLevel   int
	token         string
	hoursToExpire int
}

func New(id int64, appOrigin string, accessLevel int, hoursToExpire int) User {
	u := User{
		appOrigin:     appOrigin,
		accessLevel:   accessLevel,
		hoursToExpire: hoursToExpire,
	}
	u.SetId(id)
	return u
}
