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

type captcha struct {
	str string
}

func (c *captcha) String() string {
	return c.str
}

func Captcha(pattern, leftoperand, operator, rightoperand int) captcha {
	return captcha{
		str: "1 + one",
	}
}

func Number(n int) string {
	if n == 0 {
		return "zero"
	} else if n == 1 {
		return "one"
	} else if n == 2 {
		return "two"
	} else if n == 3 {
		return "three"
	}
}
