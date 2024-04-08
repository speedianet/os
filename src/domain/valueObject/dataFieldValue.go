package valueObject

import (
	"errors"

	voHelper "github.com/speedianet/os/src/domain/valueObject/helper"
)

type DataFieldValue string

func NewDataFieldValue(value interface{}) (DataFieldValue, error) {
	dataFieldValueStr, err := voHelper.InterfaceToString(value)
	if err != nil {
		return "", errors.New("InvalidDataFieldValue")
	}

	if len(dataFieldValueStr) <= 1 {
		return "", errors.New("DataFieldValueTooSmall")
	}

	if len(dataFieldValueStr) >= 60 {
		return "", errors.New("DataFieldValueTooBig")
	}

	return DataFieldValue(dataFieldValueStr), nil
}

func NewDataFieldValuePanic(value interface{}) DataFieldValue {
	dfv, err := NewDataFieldValue(value)
	if err != nil {
		panic(err)
	}

	return dfv
}

func (dfv DataFieldValue) String() string {
	return string(dfv)
}
