package captcha

import "testing"

// func TestNumberOneAsString(t *testing.T) {
// 	if Number(1) != "one" {
// 		t.Error("it should be one but get", Number(1))
// 	}
// }

// func TestNumberTwoAsString(t *testing.T) {
// 	if Number(2) != "two" {
// 		t.Error("it should be two but get", Number(2))
// 	}
// }

// func TestNumberAsString(t *testing.T) {
// 	var flagtest = []struct {
// 		intput int
// 		output string
// 	}{
// 		{0, "zero"},
// 		{1, "one"},
// 		{2, "two"},
// 		{3, "three"},
// 		{4, "four"},
// 		{5, "five"},
// 		{6, "six"},
// 		{7, "seven"},
// 		{8, "eight"},
// 		{9, "nine"},
// 	}
// 	for _, v := range flagtest {
// 		r := Number(v.intput)
// 		if r != v.output {
// 			t.Errorf("Expected %v from %v but got %#v", v.output, v.intput, r)
// 		}
// 	}
// }

func TestCaptchaFirstPatternOfOnePlusOne(t *testing.T) {
	c := Captcha(FirstPattern, 1, OperationPlus, 1)
	if c.String() != "1 + one" {
		t.Error("it should be 1 + one but get", c.String())
	}
}

func TestCaptchaFirstPatternOfTwoMinusOne(t *testing.T) {
	c := Captcha(FirstPattern, 2, OperationMinus, 1)
	if c.String() != "2 - one" {
		t.Error("it should be 2 - one but get", c.String())
	}
}

func TestCaptchaSecondPatternOfThreeMutiplySix(t *testing.T) {
	c := Captcha(SecondPattern, 3, OperationMutiply, 6)
	if c.String() != "three x 6" {
		t.Error("it should be three x 6 but get", c.String())
	}
}

func TestCaptchaFirstPatternOfTwoPlusOne(t *testing.T) {
	c := Captcha(FirstPattern, 2, OperationPlus, 1)
	if c.String() != "2 + one" {
		t.Error("it should be 2 + one but get", c.String())
	}
}

func TestCaptchaFirstPattern(t *testing.T) {
	testCases := []struct {
		inputPattern  int
		inputLeft     int
		inputOperator int
		inputRight    int
		expected      string
	}{
		{1, 1, 1, 1, "1 + one"},
		{1, 2, 1, 9, "2 + nine"},
		{2, 4, 2, 1, "four - 1"},
		{2, 7, 2, 1, "seven - 1"},
		{1, 2, 3, 7, "2 x seven"},
		{2, 2, 3, 9, "two x 9"},
	}
	for _, test := range testCases {
		c := Captcha(test.inputPattern, test.inputLeft, test.inputOperator, test.inputRight)
		if c.String() != test.expected {
			t.Errorf("it should be %v but got %v", test.expected, c.String())
		}
	}
}
