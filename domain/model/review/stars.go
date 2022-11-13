package review

import (
	"fmt"
)

type stars int

func newStars(value int) (*stars, error) {

	if value < 1 || value > 5 {
		err := fmt.Errorf("%s", "error arg:stars between 1 and 5")
		return nil, err
	}

	stars := stars(value)
	return &stars, nil
}
