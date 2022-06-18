package review

type publicFlg bool

func newPublicFlg(value bool) (*publicFlg, error) {

	publicFlg := publicFlg(value)
	return &publicFlg, nil
}
