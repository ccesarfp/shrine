package model

// Token Model
// id - token id
// token - user token
// **
type Token struct {
	id    string
	token string
}

func NewToken(token string) Token {
	return Token{
		token: token,
	}
}

func NewTokenWithId(id string) Token {
	return Token{
		id: id,
	}
}

func (t Token) Id() string {
	return t.id
}

func (t Token) Token() string {
	return t.token
}
