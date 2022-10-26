package vo

import (
	"fmt"
)

type UuId string

func NewUuId(value string) (*UuId, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:UuId New()")
		return nil, err
	}

	uuId := UuId(value)

	return &uuId, nil
}
