package filter

import (
	"regexp"
	"strings"
)

// 去除首尾空格
func Trim(str string) string {
	return strings.TrimSpace(str)
}

func IsPhoneInChina(phoneNumber string) bool {
	reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(phoneNumber)
}
