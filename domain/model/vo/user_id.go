package vo

import (
	"fmt"
)

type UserId string

func NewUserId(value string) (*UserId, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:userId newUserId()")
		return nil, err
	}

	userId := UserId(value)

	return &userId, nil
}
