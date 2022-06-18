package comment

import (
	"fmt"
)

type CommentId string

func NewCommentId(value string) (*CommentId, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:commentId newCommentId()")
		return nil, err
	}

	commentId := CommentId(value)

	return &commentId, nil
}
