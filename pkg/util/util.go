package util

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
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

// CreateUnixExpirationTime create expiration time
// Params:
//   - hoursToExpire: int32
//
// Return:
//   - time
//   - error
//
// **
func CreateUnixExpirationTime(hoursToExpire int32) (time.Time, error) {
	return time.Now().Add(time.Hour * time.Duration(hoursToExpire)), nil
}

// FindProcess find process
// Params:
//   - processName: string
//
// Return:
//   - *os.Process
//   - error
//
// **
func FindProcess(processName string) (*os.Process, error) {
	p := exec.Command("pgrep", processName)
	output, err := p.Output()
	if err != nil {
		return nil, err
	}

	pid := strings.TrimSpace(string(output))
	id, err := strconv.Atoi(pid)
	if err != nil {
		return nil, err
	}

	process, err := os.FindProcess(id)
	if err != nil {
		return nil, err
	}

	return process, nil
}

// SendSignal send signal to process
// Params:
//   - p: 	   *os.Process
//   - signal: os.sinal
//
// Return:
//   - bool
//   - error
//
// **
func SendSignal(p *os.Process, signal os.Signal) (bool, error) {
	err := p.Signal(signal)
	if err != nil {
		return false, err
	}
	return true, nil
}
