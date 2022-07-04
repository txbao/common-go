package utils

import (
	"encoding/json"
	"go/types"
	"strconv"
)

var UtlConvert = &_convert{}

type _convert struct {
}

func (o *_convert) SliceStringToInt(sliceString []string) (sliceInt []int) {
	if len(sliceString) == 0 {
		return
	}
	sliceInt = make([]int, len(sliceString))
	for i, str := range sliceString {
		sliceInt[i], _ = strconv.Atoi(str)
	}
	return
}

func (o *_convert) SliceInt64ToString(sliceInt []int64) (sliceString []string) {
	if len(sliceInt) == 0 {
		return
	}
	sliceString = make([]string, len(sliceInt))
	for i, num := range sliceInt {
		sliceString[i] = strconv.FormatInt(num, 10)
	}
	return
}

func (o *_convert) Int64ToString(e int64) string {
	return strconv.FormatInt(e, 10)
}

func (o *_convert) StringToInt64(e string) (int64, error) {
	return strconv.ParseInt(e, 10, 64)
}

func (o *_convert) Float64ToString(e float64) string {
	return strconv.FormatFloat(e, 'E', -1, 64)
}

//float32 转 String工具类，保留6位小数
func (o *_convert) Float32ToString(e float32) string {
	return strconv.FormatFloat(float64(e), 'f', 6, 64)
}

func (o *_convert) InterfaceToString(val interface{}) string {
	if val == nil {
		return ""
	}
	var value string
	switch val := val.(type) {
	case string:
		value = val
	case json.Number:
		value = val.String()
	case float64:
		value = o.Float64ToString(val)
	case float32:
		value = o.Float32ToString(val)
	case int:
		value = strconv.Itoa(val)
	case uint:
		value = strconv.Itoa(int(val))
	case int8:
		value = o.Int64ToString(int64(val))
	case uint8:
		value = strconv.Itoa(int(val))
	case int16:
		value = o.Int64ToString(int64(val))
	case uint16:
		value = strconv.Itoa(int(val))
	case int32:
		value = o.Int64ToString(int64(val))
	case uint32:
		value = strconv.Itoa(int(val))
	case int64:
		value = o.Int64ToString(val)
	case uint64:
		value = strconv.FormatUint(val, 10)
	case types.Nil:
		value = ""
	case []byte:
		value = string(val)
	default:
		newValue, _ := json.Marshal(val)
		value = string(newValue)
	}

	return value
}
