package captcha

import "testing"

func TestNumberOneAsString(t *testing.T) {
	if Number(1) != "one" {
		t.Error("it should be one but get", Number(1))
	}
}

func TestCaptchaFirstPatternOfOnePlusOne(t *testing.T) {
	c := Captcha(FirstPattern, 1, OperationPlus, 1)
	if c.String() != "1 + one" {
		t.Error("it should be 1 + one but get", c.String())
	}
}
