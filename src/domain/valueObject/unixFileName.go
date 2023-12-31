package valueObject

import (
	"errors"
	"regexp"
	"slices"
)

const unixFileNameRegexExpression = `^[^\n\r\t\f\0\?\[\]\<\>\/]{1,256}$`

var reservedUnixFileNames = []string{".", "..", "*", "/", "\\"}

type UnixFileName string

func NewUnixFileName(value string) (UnixFileName, error) {
	unixFileName := UnixFileName(value)
	if !unixFileName.isValid() {
		return "", errors.New("InvalidUnixFileName")
	}
	return unixFileName, nil
}

func NewUnixFileNamePanic(value string) UnixFileName {
	unixFileName, err := NewUnixFileName(value)
	if err != nil {
		panic(err)
	}
	return unixFileName
}

func (unixFileName UnixFileName) isValid() bool {
	unixFileNameRegex := regexp.MustCompile(unixFileNameRegexExpression)
	isValidFormat := unixFileNameRegex.MatchString(string(unixFileName))

	isReservedUnixFileName := slices.Contains(reservedUnixFileNames, string(unixFileName))

	return isValidFormat && !isReservedUnixFileName
}

func (unixFileName UnixFileName) String() string {
	return string(unixFileName)
}
