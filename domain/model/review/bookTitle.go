package review

import (
	"fmt"
)

type bookTitle string

func newBookTitle(value string) (*bookTitle, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:bookTitle newBookTitle()")
		return nil, err
	}

	bookTitle := bookTitle(value)
	return &bookTitle, nil
}
