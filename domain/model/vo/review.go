package vo

import(
  "fmt"
)

type Review string

func NewReview(value string) (*Review, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:Review New()")
		return nil, err
	}

	review := Review(value)

	return &review, nil
}
