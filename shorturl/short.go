package shorturl

import (
	"math"
	"math/rand"
	"regexp"
	"strings"
)

var CHARS = "InsV3Sf0obzp2i4gj1yYGqQv6wUtmBxlMAP7KHd8uTXFk9aRJWNC5EOhZDcLer"

const (
	// 0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ
	SCALE = 62
	REGEX = "^[0-9a-zA-Z]+$"
	NUM   = 6
)

func RandomStr(str string) string {
	chars := []rune(str)
	for i := len(chars) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		chars[i], chars[num] = chars[num], chars[i]
	}
	return string(chars)
}

func Encode10To62(val uint) string {
	if val < 0 {
		panic("val cannot be negative.")
	}
	str := ""
	var remainder int
	for math.Abs(float64(val)) > SCALE-1 {
		remainder = int(val % SCALE)
		str = string(CHARS[remainder]) + str
		val = val / SCALE
	}
	str = string(CHARS[val]) + str
	//for i := len(str); i < NUM; i++ {
	//	str = string(CHARS[0]) + str
	//}
	return str
}

func Decode62To10(val string) uint {
	if match, _ := regexp.MatchString(REGEX, val); !match {
		panic("input illegal.")
	}
	var result uint = 0
	index, length := 0, len(val)
	for i := 0; i < length; i++ {
		index = strings.Index(CHARS, string(val[i]))
		result += uint(index * int(math.Pow(float64(SCALE), float64(length-i-1))))
	}
	return result
}
