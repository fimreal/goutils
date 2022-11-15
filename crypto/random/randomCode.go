package random

import (
	"math/rand"
	"time"
)

const (
	NumChar      = "0123456789"
	UpAlphaChar  = "QWERTYUIOPASDFGHJKLZXCVBNM"
	LowAlphaChar = "qwertyuiopasdfghjklzxcvbnm"
	SpecChar     = "+=-@#~,.[]()$%^&*"
)

type code struct {
	// 自定义字符
	// 默认 Default:  NumChar + LowAlphaChar
	// 可选 All: "NumChar", "UpAlphaChar", "LowAlphaChar", "SpecChar"
	Composition string
	// 指定生成长度 默认 default: 8
	Length int
}

var Config *code

func init() {
	Config = &code{}
}

func (c *code) New() func() string {
	// 选定可选字符
	if len(c.Composition) == 0 {
		// 默认使用数字和小写字符组合
		c.Composition = NumChar + LowAlphaChar
	}
	// ezap.Debug("随机数生成器初始化: 确定可选字符串 ", c.Composition)

	// 确定密码长度
	if c.Length == 0 {
		c.Length = 8
	}

	rand.Seed(time.Now().UnixNano() + 19940219)
	return func() string {
		b := make([]rune, c.Length)
		compositionRune := []rune(c.Composition)
		for i := 0; i < c.Length; i++ {
			inx := rand.Intn(len(compositionRune))
			b[i] = compositionRune[inx]
		}
		return string(b)
	}
}

// 生成随机数字
func Num(length int) string {
	Config.Composition = NumChar
	Config.Length = length
	Config.New()
	return Config.New()()
}

// 生成随机字母
func Alpha(length int) string {
	Config.Composition = UpAlphaChar + LowAlphaChar
	Config.Length = length
	Config.New()
	return Config.New()()
}

// 生成随机小写字母
func LowAlpha(length int) string {
	Config.Composition = LowAlphaChar
	Config.Length = length
	Config.New()
	return Config.New()()
}

// 生成随机大写字母
func UpAlpha(length int) string {
	Config.Composition = UpAlphaChar
	Config.Length = length
	Config.New()
	return Config.New()()
}

// 生成随机密码，小写字母和数字组合
func GenPassword(length int) string {
	Config = &code{}
	Config.Length = length
	Config.New()
	return Config.New()()
}
