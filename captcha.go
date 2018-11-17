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
	numString := ""
	if n == 0 {
		numString = "zero"
	} else if n == 1 {
		numString = "one"
	} else if n == 2 {
		numString = "two"
	} else if n == 3 {
		numString = "three"
	} else if n == 4 {
		numString = "four"
	} else if n == 5 {
		numString = "five"
	} else if n == 6 {
		numString = "six"
	} else if n == 7 {
		numString = "seven"
	} else if n == 8 {
		numString = "eight"
	} else if n == 9 {
		numString = "nine"
	}
	return numString
}
