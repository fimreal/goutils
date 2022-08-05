package random

import (
	"testing"
)

func TestNew(t *testing.T) {
	c := &code{}
	randomCode := c.New()
	t.Log(randomCode())
	t.Log(randomCode())

	c = &code{
		Composition: NumChar + LowAlphaChar + UpAlphaChar + SpecChar,
		Length:      16,
	}
	randomCode = c.New()
	t.Log(randomCode())
	t.Log(randomCode())

	c = &code{
		Length:      32,
		Composition: NumChar + LowAlphaChar + "∂ƒœ´",
	}
	randomCode = c.New()
	t.Log(randomCode())
	t.Log(randomCode())
}

func TestNum(t *testing.T) {
	t.Log("8num: ", Num(8))
}

func TestAlpha(t *testing.T) {
	t.Log("8 alpha: ", Alpha(8))
}

func TestLowAlpha(t *testing.T) {
	t.Log("8 lowalpha: ", LowAlpha(8))
}

func TestUpAlpha(t *testing.T) {
	t.Log("8 upalpha: ", UpAlpha(8))
}

func TestGenpassword(t *testing.T) {
	t.Log("8 password: ", GenPassword(8))
}
