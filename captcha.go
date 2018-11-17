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

func Number(n int) string {
<<<<<<< HEAD
	return "two"
=======
	return "one"
>>>>>>> 70fff1e81d57ddb87a0928de7b59cc2dae758c56
}
