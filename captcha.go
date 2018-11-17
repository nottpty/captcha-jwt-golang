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
	OperationDivide
)

type captcha struct {
	pattern, leftOperand, operator, rightOperand int
}

func (c *captcha) String() string {
	leftStr := strconv.Itoa(c.leftOperand)
	rightStr := strconv.Itoa(c.rightOperand)
	var operatorStr string
	if c.operator == OperationPlus {
		operatorStr = "+"
	}
	if c.operator == OperationMinus {
		operatorStr = "-"
	}
	if c.operator == OperationMutiply {
		operatorStr = "x"
	}
	if c.operator == OperationDivide {
		operatorStr = "/"
	}
	if c.pattern == FirstPattern {
		return leftStr + " " + operatorStr + " " + Number(c.rightOperand)
	} else {
		return Number(c.leftOperand) + " " + operatorStr + " " + rightStr
	}

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
