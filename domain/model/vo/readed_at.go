package vo

import(
	"time"
)

type ReadedAt time.Time

func NewReadedAt(value time.Time) (*ReadedAt, error) {
	readedAt := ReadedAt(value)
	return &readedAt, nil
}
