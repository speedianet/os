package valueObject

import (
	"errors"
	"strconv"

	voHelper "github.com/speedianet/os/src/domain/valueObject/helper"
)

type AccountId uint64

func NewAccountId(value interface{}) (AccountId, error) {
	accountId, err := voHelper.InterfaceToUint64(value)
	if err != nil {
		return 0, errors.New("AccountIdMustBeInt")
	}

	return AccountId(accountId), nil
}

func (vo AccountId) Get() uint64 {
	return uint64(vo)
}

func (vo AccountId) String() string {
	return strconv.FormatUint(uint64(vo), 10)
}
