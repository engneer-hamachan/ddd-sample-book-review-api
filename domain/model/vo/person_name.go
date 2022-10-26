package vo

import(
  "fmt"
)

type PersonName string

func NewPersonName(value string) (*PersonName, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:PersonName New()")
		return nil, err
	}

	personName := PersonName(value)

	return &personName, nil
}
