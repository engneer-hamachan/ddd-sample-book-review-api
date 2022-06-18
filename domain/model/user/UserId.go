package user

import (
	"fmt"
)

type UserId string

func NewUserId(value string) (*UserId, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:userId NewUserId()")
		return nil, err
	}

	userId := UserId(value)

	return &userId, nil
}
