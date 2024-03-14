package circuit_open

import (
	"fmt"
)

type Error struct {
}

func (e *Error) Error() string {
	return fmt.Sprintf("the application is not available")
}
