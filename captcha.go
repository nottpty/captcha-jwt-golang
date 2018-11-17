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
		if rightoperand == One {
			rightStr = "one"
		}
		if rightoperand == Two {
			rightStr = "two"
		}
	}
	result := leftStr + " " + operatorStr + " " + rightStr
	return captcha{
		str: result,
	}
}

func Number(n int) string {
	numString := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	return numString[n]
}
