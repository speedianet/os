package valueObject

import (
	"errors"
	"regexp"
)

const unixFilePermissionsRegexExpression = `^[0-7]{3,4}$`

type UnixFilePermissions string

func NewUnixFilePermissions(value string) (UnixFilePermissions, error) {
	unixFilePermissions := UnixFilePermissions(value)
	if !unixFilePermissions.isValid() {
		return "", errors.New("InvalidUnixFilePermissions")
	}

	return unixFilePermissions, nil
}

func NewUnixFilePermissionsPanic(value string) UnixFilePermissions {
	unixFilePermissions, err := NewUnixFilePermissions(value)
	if err != nil {
		panic(err)
	}
	return UnixFilePermissions(unixFilePermissions)
}

func (unixFilePermissions UnixFilePermissions) isValid() bool {
	unixFilePermissionsRegex := regexp.MustCompile(unixFilePermissionsRegexExpression)
	return unixFilePermissionsRegex.MatchString(string(unixFilePermissions))
}

func (unixFilePermission UnixFilePermissions) String() string {
	return string(unixFilePermission)
}
