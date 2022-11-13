package review

import (
	"fmt"
)

type reviewTitle string

func newReviewTitle(value string) (*reviewTitle, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:reviewTitle newReviewTitle()")
		return nil, err
	}

	reviewTitle := reviewTitle(value)
	return &reviewTitle, nil
}
