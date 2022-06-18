package review_like

import (
	"fmt"
)

type reviewLikeId string

func newReviewLikeId(value string) (*reviewLikeId, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:reviewLikeId newReviewLikeId()")
		return nil, err
	}

	reviewLikeId := reviewLikeId(value)

	return &reviewLikeId, nil
}
