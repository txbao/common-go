package utils

import (
	"math/rand"
	"strconv"
	"time"
)

var UtlUnique = &_unique{}

type _unique struct {
}

func init() {
	rand.Seed(UtlUnique.Number())
}

// 19位的唯一数字
func (o *_unique) Number() int64 {
	nano1 := time.Now().UnixNano()
	time.Sleep(time.Microsecond)
	nano2 := time.Now().UnixNano()
	return nano1 + nano2/100%100
}

// 12位的唯一字符串
func (o *_unique) String() string {
	return strconv.FormatInt(o.Number(), 36)
}

// 任意位数的唯一字符串
func (o *_unique) StringEx(strLen int) string {
	return o.String() + o.RandStringLower(strLen)
}

// 随机字符串
func (o *_unique) RandString(strLen int) string {
	var c int32
	var cList []int32
	for i := 0; i < strLen; i++ {
		c = rand.Int31n(62)
		if c < 10 {
			c += 48
		} else if c < 36 {
			c += 55
		} else {
			c += 61
		}
		cList = append(cList, c)
	}
	return string(cList)
}

// 随机字符串-仅小写
func (o *_unique) RandStringLower(strLen int) string {
	var c int32
	var cList []int32
	for i := 0; i < strLen; i++ {
		c = rand.Int31n(36)
		if c < 10 {
			c += 48
		} else {
			c += 87
		}
		cList = append(cList, c)
	}
	return string(cList)
}
