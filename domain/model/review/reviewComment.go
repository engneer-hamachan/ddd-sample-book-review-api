package review

import (
	"fmt"
)

type reviewComment string

func newReviewComment(value string) (*reviewComment, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:review newReviewComment()")
		return nil, err
	}

	review := reviewComment(value)
	return &review, nil
}
