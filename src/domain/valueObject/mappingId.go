package valueObject

import (
	"errors"
	"strconv"

	voHelper "github.com/speedianet/os/src/domain/valueObject/helper"
)

type MappingId uint64

func NewMappingId(value interface{}) (MappingId, error) {
	mId, err := voHelper.InterfaceToUint(value)
	if err != nil {
		return 0, errors.New("InvalidMappingId")
	}

	return MappingId(mId), nil
}

func NewMappingIdPanic(value interface{}) MappingId {
	mId, err := NewMappingId(value)
	if err != nil {
		panic(err)
	}
	return mId
}

func (mId MappingId) Get() uint64 {
	return uint64(mId)
}

func (mId MappingId) String() string {
	return strconv.FormatUint(uint64(mId), 10)
}
