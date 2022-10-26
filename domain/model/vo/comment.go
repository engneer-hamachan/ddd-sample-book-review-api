package vo

import(
  "fmt"
)

type Comment string

func NewComment(value string) (*Comment, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:Comment New()")
		return nil, err
	}

	comment := Comment(value)

	return &comment, nil
}
