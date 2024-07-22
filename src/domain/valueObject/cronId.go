package valueObject

import (
	"errors"
	"strconv"

	voHelper "github.com/speedianet/os/src/domain/valueObject/helper"
)

type CronId uint64

func NewCronId(value interface{}) (cronId CronId, err error) {
	uintValue, err := voHelper.InterfaceToUint64(value)
	if err != nil {
		return cronId, errors.New("CronIdMustBeInt")
	}

	return CronId(uintValue), nil
}

func (vo CronId) Uint() uint64 {
	return uint64(vo)
}

func (vo CronId) String() string {
	return strconv.FormatUint(uint64(vo), 10)
}
