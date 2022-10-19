package str

import (
	"strings"
)

/*
*
获取字符串长度
*/
func StringCount(str string) int {
	num := strings.Count(str, "") - 1
	return num
}

// HideString 给字段加***
func HideString(src string, hLen int) string {
	str := []rune(src)
	if hLen == 0 {
		hLen = 4
	}
	hideStr := ""
	for i := 0; i < hLen; i++ {
		hideStr += "*"
	}
	hideLen := len(str) / 2
	showLen := len(str) - hideLen
	if hideLen == 0 || showLen == 0 {
		return hideStr
	}
	subLen := showLen / 2
	if subLen == 0 {
		return string(str[:showLen]) + hideStr
	}
	s := string(str[:subLen])
	s = s + hideStr
	s = s + string(str[len(str)-subLen:])
	return s
}

func Int32ToString(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}
