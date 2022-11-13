package review

import (
	"fmt"
)

type publisher string

func newPublisher(value string) (*publisher, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:publisher newPublisher()")
		return nil, err
	}

	publisher := publisher(value)
	return &publisher, nil
}
