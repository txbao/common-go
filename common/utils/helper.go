package utils

/**
txbao
*/

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

func SmartPrint(i interface{}) {
	var kv = make(map[string]interface{})
	vValue := reflect.ValueOf(i)
	vType := reflect.TypeOf(i)
	for i := 0; i < vValue.NumField(); i++ {
		kv[vType.Field(i).Name] = vValue.Field(i)
	}
	fmt.Println("获取到数据:")
	for k, v := range kv {
		fmt.Print(k)
		fmt.Print(":")
		fmt.Print(v)
		fmt.Println()
	}
}

/**
 * 生成订单号
 * @param int unique 唯一标识，建议用户ID
 * @return string
 */
func GenerateOrderNo(unique string) string {
	t := time.Now() //2019-07-31 13:55:21.3410012 +0800 CST m=+0.006015601
	orderNo := unique + t.Format("20060102150405")

	nano := t.Nanosecond()
	nano_str := Int64ToString(int64(nano))
	if len(nano_str) > 6 {
		orderNo += nano_str[0:6]
	}
	orderNo += RandomString(4, 1)

	return orderNo
}

//生成随机数  n为倍数 t=1为数字 t=2为数字+字母
func RandomString(n int, t int) string {
	var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var NumberLetters = []rune("0123456789")
	var letters []rune

	if t == 1 {
		letters = NumberLetters
	} else {
		letters = defaultLetters
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	//当为数字时，首为不为0
	if t == 1 {
		bstr := string(b)
		if bstr != "" {
			if bstr[0:1] == "0" {
				randn := rand.Intn(8) + 1
				return strconv.Itoa(randn) + bstr[1:n]
			}
		}
	}

	return string(b)
}

func RandStr(strLen int) string {
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

//生成8位随机数
func CreateCaptcha() string {
	return fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000000))
}

//生成兑换码
func CreateCode(total int, digit int) []string {
	var codeArr []string
	//var b58Alphabet = []byte("123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")
	var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZ")
	t := time.Now()
	rand.Seed(time.Now().Unix())
	m := make(map[string]struct{})
	//total := 1000000
	a := 0
	for {
		id := make([]byte, digit)
		for i := 0; i < digit; i++ {
			id[i] = b58Alphabet[rand.Int()%len(b58Alphabet)]
		}
		idstr := string(id)
		// 去除重复
		if _, ok := m[idstr]; ok {
			continue
		}
		m[idstr] = struct{}{}
		a++
		if a >= total {
			break
		}

		codeArr = append(codeArr, idstr)
	}
	result := len(m)
	fmt.Println("生成", total)
	fmt.Printf("重复率%f\n", float64(total-result)/float64(total))
	fmt.Println("用时", time.Since(t))
	return codeArr
}

//struct转map，用于搜索
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj) // 获取 obj 的类型信息
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr { // 如果是指针，则获取其所指向的元素
		t = t.Elem()
		v = v.Elem()
	}

	var data = make(map[string]interface{})
	if t.Kind() == reflect.Struct { // 只有结构体可以获取其字段信息
		for i := 0; i < t.NumField(); i++ {
			data[t.Field(i).Name] = v.Field(i).Interface()
		}

	}
	return data
}

//获取MapKey
func MapValue(resMap map[string]interface{}, key string) string {
	if value, ok := resMap[key]; ok {
		return Interface2string(value)
	}
	return ""
}

//map合并
func MapMerge(map1 map[string]interface{}, map2 map[string]interface{}) map[string]interface{} {
	//result := []map[string]interface{}{}

	mp3 := make(map[string]interface{})
	for k, v := range map1 {
		if _, ok := map1[k]; ok {
			mp3[k] = v
		}
	}

	for k, v := range map2 {
		if _, ok := map2[k]; ok {
			mp3[k] = v
		}
	}
	return mp3
	//result = append(result, map1, map2)
	//fmt.Println(result)
}

//截取银行卡后四位
func BankCodeLast(cardNo string) string {
	//截取银行卡后四位
	if len(cardNo) <= 4 {
		return ""
	}
	cardNoLen := len(cardNo)
	return cardNo[cardNoLen-4 : cardNoLen]
}

//截取字符串后几位
func StringGetLastChar(str string, digit int) string {
	//截取后四位
	if len(str) <= digit {
		return ""
	}
	cardNoLen := len(str)
	return str[cardNoLen-digit : cardNoLen]
}
