package fun

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	guuid "github.com/google/uuid"
)

func ArrayKey(data []map[string]interface{}, key string) (valData string) {
	for i := 0; i < len(data); i++ {
		v := data[i]
		if val, ok := v[key]; ok {
			if vd, iok := val.(string); iok {
				valData = vd
				break
			}
		}
	}
	return valData
}

/*
***
key to arr find val ==
*/
func InArray(key string, arr []string) bool {
	arrNum := len(arr)
	if arrNum > 0 {
		for i := 0; i < arrNum; i++ {
			val := arr[i]
			if key == val {
				return true
			}
		}
	}
	return false
}

func Explode(val, key string) (arr []string) {
	arr = strings.Split(val, key)
	return arr
}

/*
*
CreateUuid
*/
func CreateUuid() string {
	id := guuid.New()
	return id.String()
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func CreateRandNumber(number int) string {
	numberString := strconv.Itoa(number)
	return fmt.Sprintf("%0"+numberString+"v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

func CreateRandNumberInt32() int32 {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := rnd.Int31n(99999999)
	return code
}

func ToInt(num string) int {
	number, _ := strconv.Atoi(num)
	return number
}

func GetSumBytes(bytes int64) int64 {
	var (
		kb  int64 = 1024
		mb  int64 = 0
		gb  int64 = 0
		tb  int64 = 0
		sum int64 = 0
	)

	mb = kb * 1024
	gb = mb * 1024
	tb = gb * 1024

	if bytes < kb {
		sum = bytes
	} else if bytes < mb {
		sum = bytes / kb
	} else if bytes < gb {
		sum = bytes / mb
	} else if bytes < tb {
		sum = bytes / gb
	}

	return sum
}

func GetBytes(number int64, inType string) int64 {
	var sum int64 = 0

	if inType == "kb" {
		sum = number * 1024
	} else if inType == "mb" {
		sum = (1024 * 1024) * number
	} else if inType == "gb" {
		sum = (1024 * 1024 * 1024) * number
	} else if inType == "tb" {
		sum = (1024 * 1024 * 1024 * 1024) * number
	}
	return sum
}

// ???????????????
func SubString(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	return string(r[start:end])
}

func RemoveRepByLoop(slc []int) []int {
	result := []int{} // ????????????
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // ??????????????????????????????false
				break
			}
		}
		if flag { // ?????????false?????????????????????
			result = append(result, slc[i])
		}
	}
	return result
}

// ??????map???????????????????????????????????????
func RemoveRepByMap(slc []int) []int {
	result := []int{}
	tempMap := map[int]byte{} // ?????????????????????
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // ??????map??????map?????????????????????????????????
			result = append(result, e)
		}
	}
	return result
}

func RemoveRep(slc []int) []int {
	if len(slc) < 1024 {
		// ??????????????????1024???????????????????????????
		return RemoveRepByLoop(slc)
	} else {
		// ????????????????????????map?????????
		return RemoveRepByMap(slc)
	}
}

func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}
