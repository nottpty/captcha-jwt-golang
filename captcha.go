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
	pattern, leftOperand, operator, rightOperand int
}

func (c *captcha) String() string {
	leftStr := strconv.Itoa(c.leftOperand)
	var operatorStr string
	if c.pattern == OperationPlus {
		if c.operator == OperationPlus {
			operatorStr = "+"
		}
		if c.operator == OperationMinus {
			operatorStr = "-"
		}

	}
	return leftStr + " " + operatorStr + " " + Number(c.rightOperand)
}

func Captcha(pattern, leftOperand, operator, rightOperand int) captcha {

	return captcha{
		pattern:      pattern,
		leftOperand:  leftOperand,
		operator:     operator,
		rightOperand: rightOperand,
	}
}

func Number(n int) string {
	numString := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	return numString[n]
}
