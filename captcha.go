package captcha

const (
	FirstPattern = iota + 1
	SecondPattern
)

const (
	OperationPlus = iota + 1
	OperationMinus
	OperationMutiply
)

type captcha struct{}

func Captcha(pattern, leftoperand, operator, rightoperand int) captcha {
	return captcha{}
}

// Number change integer to string
func Number() string {

	return ""
}
