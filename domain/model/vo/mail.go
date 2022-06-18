package vo

import (
	"fmt"
)

type Mail string

func NewMail(value string) (*Mail, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:mail NewMail()")
		return nil, err
	}

	mail := Mail(value)
	return &mail, nil
}
