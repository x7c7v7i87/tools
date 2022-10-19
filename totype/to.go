package totype

import (
	"fmt"
	"strconv"
	//"errors"
)

type To struct {
	val interface{}
}

func SetVal(val interface{}) *To {
	return &To{val: val}
}

func (t *To) String() string {
	val := t.val
	if vd, iok := val.(string); iok {
		return vd
	}
	return ""
}

func (t *To) StringToInt() (int, error) {
	str := t.String()
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

func (t *To) StringInInt32() (int32, error) {
	str := t.String()

	i32, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(i32), nil
}

func (t *To) StringInInt64() (int64, error) {
	str := t.String()
	i64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return int64(i64), nil
}

func (t *To) Int() int {
	val := t.val
	if vd, iok := val.(int); iok {
		return vd
	}
	return 0
}

func (t *To) IntInString() string {
	var i int
	val := t.val
	if vd, iok := val.(int); iok {
		i = vd
	}
	if i > 0 {
		return strconv.Itoa(i)
	} else {
		return ""
	}
}

func (t *To) Int32InString() string {
	return fmt.Sprint(t.val)
}

func (t *To) Kv() (vd map[string]interface{}) {
	val := t.val
	if vd, ok := val.(map[string]interface{}); ok {
		return vd
	}
	return
}
