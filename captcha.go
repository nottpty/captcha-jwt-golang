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

<<<<<<< HEAD
// Number change integer to string
func Number() string {

	return ""
=======
func Number(n int) string {
	return "one"
>>>>>>> 690bdf2d0bee307b30575e146e683cc854d8b03c
}
