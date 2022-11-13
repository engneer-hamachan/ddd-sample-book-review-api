package review

import (
	"time"
)

type readedAt time.Time

func newReadedAt(value time.Time) (*readedAt, error) {
	readedAt := readedAt(value)
	return &readedAt, nil
}
