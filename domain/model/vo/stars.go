package vo

import(
	"fmt"
)

type Stars int

func NewStars(value int) (*Stars, error) {
	if value < 1 || value > 5 {
		err := fmt.Errorf("%s", "error arg:stars between 1 and 5")
		return nil, err
	}

	stars := Stars(value)
	return &stars, nil
}
