package vo

import(
  "fmt"
)

type Publisher string

func NewPublisher(value string) (*Publisher, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:Publisher New()")
		return nil, err
	}

	publisher := Publisher(value)

	return &publisher, nil
}
