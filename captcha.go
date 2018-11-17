package captcha

import (
	"fmt"
	"strconv"
)

const (
	FirstPattern = iota + 1
	SecondPattern
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

var operators = map[int]string{
	1: "+",
	2: "-",
	3: "x",
}

func (c *captcha) String() string {
	leftStr := strconv.Itoa(c.leftOperand)
	rightStr := strconv.Itoa(c.rightOperand)

	if c.pattern == FirstPattern {
		return fmt.Sprintf("%s %s %s", leftStr, operators[c.operator], Number(c.rightOperand))
	}

	return fmt.Sprintf("%s %s %s", Number(c.leftOperand), operators[c.operator], rightStr)
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
