package util

import (
	"fmt"
	"strings"
)

// PrepareKey Prepare key
// Params:
//   - id: user id
//   - appName: data source application
//
// Return:
//   - formatted key with user id + application name
//
// **
func PrepareKey(id int, appName string) string {
	return fmt.Sprintf("%d-%s", id, strings.ToLower(appName))
}
