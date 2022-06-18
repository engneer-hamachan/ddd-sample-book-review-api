package user

import (
	"fmt"
)

type name string

func newName(value string) (*name, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:name newName()")
		return nil, err
	}

	name := name(value)
	return &name, nil
}
