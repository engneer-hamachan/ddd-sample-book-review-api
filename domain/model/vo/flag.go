package vo

type Flag bool

func NewFlag(value bool) (*Flag, error) {
	flag := Flag(value)
	return &flag, nil
}
