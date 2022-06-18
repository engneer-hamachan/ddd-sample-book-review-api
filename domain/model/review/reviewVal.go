package review

import (
	"fmt"
)

type reviewVal string

func newReviewComment(value string) (*reviewVal, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:review newReviewComment()")
		return nil, err
	}

	review := reviewVal(value)
	return &review, nil
}
