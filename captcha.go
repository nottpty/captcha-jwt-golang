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
<<<<<<< HEAD
	return "two"
=======
	return "one"
>>>>>>> 70fff1e81d57ddb87a0928de7b59cc2dae758c56
}
