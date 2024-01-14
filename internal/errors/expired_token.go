package errors

import "fmt"

type ExpiredToken struct {
}

func (e *ExpiredToken) Error() string {
	return fmt.Sprintf("the token has expired")
}
