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
