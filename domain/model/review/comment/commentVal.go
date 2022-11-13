package comment

import (
	"fmt"
)

type commentVal string

func newCommentVal(value string) (*commentVal, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:comment newCommentVal()")
		return nil, err
	}

	comment := commentVal(value)
	return &comment, nil
}
