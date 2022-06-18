package vo

import (
	"fmt"
)

type UserName string

func NewUserName(value string) (*UserName, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:name NewUserName()")
		return nil, err
	}

	name := UserName(value)
	return &name, nil
}
