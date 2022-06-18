package comment_like

import (
	"fmt"
)

type commentLikeId string

func newCommentLikeId(value string) (*commentLikeId, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:commentLikeId newCommentLikeId()")
		return nil, err
	}

	commentLikeId := commentLikeId(value)

	return &commentLikeId, nil
}
