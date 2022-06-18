package user

import (
	"fmt"
)

type password string

func newPassword(value string) (*password, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:password NewPassword()")
		return nil, err
	}

	password := password(value)
	return &password, nil
}
