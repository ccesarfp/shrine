package util

import (
	"fmt"
	"regexp"
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
func PrepareKey(id int64, appName string) string {
	return fmt.Sprintf("%d-%s", id, strings.ToLower(appName))
}

// ValidateUsingRegex Validate value with regex
// Params:
//   - pattern: regex
//   - value: value to validate
//
// Return:
//   - true or false
//
// **
func ValidateUsingRegex(pattern string, value string) (bool, error) {
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return false, err
	}
	isValid := regex.MatchString(value)

	return isValid, nil
}
