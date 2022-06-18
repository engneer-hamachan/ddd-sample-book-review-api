package user

import (
	"fmt"
)

type userId string

func newUserId(value string) (*userId, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:userId newUserId()")
		return nil, err
	}

	userId := userId(value)

	return &userId, nil
}
