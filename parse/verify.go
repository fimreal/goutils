package parse

import (
	"regexp"

	n "github.com/fimreal/goutils/net"
)

func IsEmail(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{2,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}

// 国内手机号
func IsPhoneNumber(phone string) bool {
	pattern := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	re := regexp.MustCompile(pattern)
	return re.MatchString(phone)
}

func IsUsername(username string) bool {
	pattern := `^[a-z][_.-a-z0-9]{2,31}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(username)
}

func IsIPv4(ip string) bool {
	return n.IsIPv4(ip)
}

func IsIPv6(ip string) bool {
	return n.IsIPv6(ip)
}
