package valueObject

import (
	"errors"
	"regexp"
)

const mimeTypeRegexExpression = `^[\p{L}0-9\-]{1,64}\/[\p{L}0-9\-\_\+\.\,]{2,128}$|^(directory|generic)$`

type MimeType string

func NewMimeType(value string) (MimeType, error) {
	mimeType := MimeType(value)
	if !mimeType.isValid() {
		return "", errors.New("InvalidMimeType")
	}
	return mimeType, nil
}

func NewMimeTypePanic(value string) MimeType {
	mimeType, err := NewMimeType(value)
	if err != nil {
		panic(err)
	}
	return mimeType
}

func (mimeType MimeType) isValid() bool {
	mimeTypeRegex := regexp.MustCompile(mimeTypeRegexExpression)
	return mimeTypeRegex.MatchString(string(mimeType))
}

func (mimeType MimeType) String() string {
	return string(mimeType)
}
