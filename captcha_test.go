package captcha

import "testing"

func TestCaptchaFirstPatternOfOnePlusOne(t *testing.T) {
	c := Captcha(FirstPattern, 1, OperationPlus, 1)
	if c.String() != "1 + one" {
		t.Error("it should be 1 + one but get", c.String())
	}
}
