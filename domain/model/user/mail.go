package user

import (
	"fmt"
)

type mail string

func newMail(value string) (*mail, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:mail newMail()")
		return nil, err
	}

	mail := mail(value)
	return &mail, nil
}
