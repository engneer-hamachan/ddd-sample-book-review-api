package vo

import(
  "fmt"
)

type Title string

func NewTitle(value string) (*Title, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:Title New()")
		return nil, err
	}

	title := Title(value)

	return &title, nil
}
