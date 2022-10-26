package vo

import (
	"fmt"
)

type Email string

func NewEmail(value string) (*Email, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:Email New()")
		return nil, err
	}

	email := Email(value)

	return &email, nil
}
