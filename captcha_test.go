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

func TestNumberAsString(t *testing.T) {
	var flagtest = []struct {
		intput int
		output string
	}{
		{0, "zero"},
		{1, "one"},
		{2, "two"},
		{3, "three"},
	}
	for _, v := range flagtest {
		r := Number(v.intput)
		if r != v.output {
			t.Errorf("Expected %v from %v but got %v", v.output, v.intput, r)
		}
	}
}

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
	if c.String() != "Three x 6" {
		t.Error("it should be Three x  but get", c.String())
	}
}

func TestCaptchaFirstPatternOfTwoPlusOne(t *testing.T) {
	c := Captcha(FirstPattern, 2, OperationPlus, 1)
	if c.String() != "2 + one" {
		t.Error("it should be 2 + one but get", c.String())
	}
}
