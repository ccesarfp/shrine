package expired_token

import "fmt"

type Error struct {
}

func (e *Error) Error() string {
	return fmt.Sprintf("the token has expired")
}
