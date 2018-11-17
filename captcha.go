package captcha

import "strconv"

const (
	FirstPattern = iota + 1
	SecondPattern
)

const (
	Zero = iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
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
	leftStr := strconv.Itoa(leftoperand)
	rightStr := strconv.Itoa(rightoperand)
	var operatorStr string
	if pattern == OperationPlus {
		if operator == OperationPlus {
			operatorStr = "+"
		}
		if operator == OperationMinus {
			operatorStr = "-"
		}

	}
	result := leftStr + " " + operatorStr + " " + Number(rightoperand)
	return captcha{
		str: result,
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
	return ""
}
