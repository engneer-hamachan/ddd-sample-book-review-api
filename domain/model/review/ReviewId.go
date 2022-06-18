package review

import (
	"fmt"
)

type ReviewId string

func NewReviewId(value string) (*ReviewId, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:reviewId NewReviewId()")
		return nil, err
	}

	reviewId := ReviewId(value)

	return &reviewId, nil
}
