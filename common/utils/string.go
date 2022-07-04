/**
 * 字符串操作
 * Create by whimp(whimp@189.cn)
 * Date: 2020\5\8 14:51
 */
package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"go/types"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

func StringToInt64(e string) (int64, error) {
	return strconv.ParseInt(e, 10, 64)
}

func IntToString(e int) string {
	return strconv.Itoa(e)
}

func Float64ToString(e float64) string {
	return strconv.FormatFloat(e, 'f', -1, 64)
}

//float32 转 String工具类，保留6位小数
func Float32ToString(input_num float32) string {
	// to convert a float number to a string
	return strconv.FormatFloat(float64(input_num), 'f', 6, 64)
}

//txbao
func StringToFloat64(e string) float64 {
	value, _ := strconv.ParseFloat(e, 64)
	return value
}

//txbao
func Int64ToInt(int64_num int64) int {
	// 将 int64 转化为 int
	int_num := *(*int)(unsafe.Pointer(&int64_num))
	return int_num
}

func Int64ToString(e int64) string {
	return strconv.FormatInt(e, 10)
}

//10进制转16 txbao
func IntHex10to16(ten int) string {
	m := 0
	hex := make([]int, 0)
	for {
		m = ten % 16
		ten = ten / 16
		if ten == 0 {
			hex = append(hex, m)
			break
		}
		hex = append(hex, m)
	}
	hexStr := []string{}
	for i := len(hex) - 1; i >= 0; i-- {
		if hex[i] >= 10 {
			hexStr = append(hexStr, fmt.Sprintf("%c", 'A'+hex[i]-10))
		} else {
			hexStr = append(hexStr, fmt.Sprintf("%d", hex[i]))
		}
	}
	return strings.Join(hexStr, "")
}

//转为整数，类似于PHP的 Intval txbao
func Intval(someString string) string {
	chiReg := regexp.MustCompile("[^0-9]")
	return chiReg.ReplaceAllString(someString, "")
}

//获取URL中批量id并解析
func IdsStrToIdsInt64Group(key string, c *gin.Context) []int64 {
	IDS := make([]int64, 0)
	ids := strings.Split(c.Param(key), ",")
	for i := 0; i < len(ids); i++ {
		ID, _ := strconv.ParseInt(ids[i], 10, 64)
		IDS = append(IDS, ID)
	}
	return IDS
}

func GetCurrntTime() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

func GetLocation(ip string) string {
	if ip == "127.0.0.1" || ip == "localhost" {
		return "内部IP"
	}
	resp, err := http.Get("https://restapi.amap.com/v3/ip?ip=" + ip + "&key=3fabc36c20379fbb9300c79b19d5d05e")
	if err != nil {
		panic(err)

	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(s))

	m := make(map[string]string)

	err = json.Unmarshal(s, &m)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
	}
	if m["province"] == "" {
		return "未知位置"
	}
	return m["province"] + "-" + m["city"]
}

func StructToJsonStr(e interface{}) (string, error) {
	if b, err := json.Marshal(e); err == nil {
		return string(b), err
	} else {
		return "", err
	}
}

func GetBodyString(c *gin.Context) (string, error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return string(body), nil
	} else {
		return "", err
	}
}

func JsonStrToMap(e string) (map[string]interface{}, error) {
	var dict map[string]interface{}
	if err := json.Unmarshal([]byte(e), &dict); err == nil {
		return dict, err
	} else {
		return nil, err
	}
}
func StrToInt(err error, index string) int {
	result, err := strconv.Atoi(index)
	if err != nil {
		AssertErr(err, "string to int error"+err.Error(), -1)
	}
	return result
}

/**
字符转为整数 txbao
*/
func StringToInt(v string) int {
	if len(v) == 0 {
		return 0
	}
	if i, err := strconv.Atoi(v); err == nil {
		return i
		//fmt.Printf("%T, %v", s, s)
	} else {
		AssertErr(err, "string to int error"+err.Error(), -1)
	}
	return 0
}

//加密
func Encrypt(e string) ([]byte, error) {
	//var s []byte
	s, err := bcrypt.GenerateFromPassword([]byte(e), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return s, err
}

func CompareHashAndPassword(e string, p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		log.Print(err.Error())
		return false, err
	}
	return true, nil
}

// Assert 条件断言
// 当断言条件为 假 时触发 panic
// 对于当前请求不会再执行接下来的代码，并且返回指定格式的错误信息和错误码
func Assert(condition bool, msg string, code ...int) {
	if !condition {
		statusCode := 200
		if len(code) > 0 {
			statusCode = code[0]
		}
		panic("CustomErroe#" + strconv.Itoa(statusCode) + "#" + msg)
	}
}

// AssertErr 错误断言
// 当 error 不为 nil 时触发 panic
// 对于当前请求不会再执行接下来的代码，并且返回指定格式的错误信息和错误码
// 若 msg 为空，则默认为 error 中的内容
func AssertErr(err error, msg string, code ...int) {
	if err != nil {
		statusCode := 200
		if len(code) > 0 {
			statusCode = code[0]
		}
		if msg == "" {
			msg = err.Error()
		}
		panic("CustomError#" + strconv.Itoa(statusCode) + "#" + msg)
	}
}

// interface转string txbao
func Interface2string(val interface{}) string {
	fmt.Println("reflect", reflect.TypeOf(val))
	//fmt.Println(fmt.Sprintf("%T", val))
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
		value = Float64ToString(val)
	case float32:
		value = Float32ToString(val)
	case int:
		value = IntToString(val)
	case uint:
		value = strconv.Itoa(int(val))
	case int8:
		value = Int64ToString(int64(val))
	case uint8:
		value = strconv.Itoa(int(val))
	case int16:
		value = Int64ToString(int64(val))
	case uint16:
		value = strconv.Itoa(int(val))
	case int32:
		value = Int64ToString(int64(val))
	case uint32:
		value = strconv.Itoa(int(val))
	case int64:
		value = Int64ToString(val)
	case uint64:
		value = strconv.FormatUint(val, 10)
	case types.Nil:
		value = ""
	case []byte:
		value = string(val)
	case *resty.Response:
		value = val.String()
	default:
		newValue, _ := json.Marshal(val)
		value = string(newValue)
	}

	return value
}

//限制空格+换行
func TrimStr(str string) string {
	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	return str
}

//string hex convert to big.Int
func HexToBigInt(hex string) *big.Int {
	n := new(big.Int)
	n, _ = n.SetString(hex[2:], 16)

	return n
}

//string big.Int convert to  hex
func BigIntToHex(bigInt *big.Int) string {
	count := new(big.Int)
	count.SetString(bigInt.String(), 10)
	return Strtoupper(ToHexInt(count))
}

func ToHexInt(n *big.Int) string {
	return fmt.Sprintf("%x", n) // or %X or upper case
}

//格式化金额，去除多余的字母及非法金额
func MoneyFormat(money string) string {
	reg, err := regexp.Compile("[^0-9.]+")
	if err != nil {
		fmt.Println(err)
	}
	return reg.ReplaceAllString(money, "")
}

//map[string]interface{} 转 map[string]string
func MapInteface2string(m map[string]interface{}) map[string]string {
	ret := make(map[string]string, len(m))
	for k, v := range m {
		ret[k] = fmt.Sprint(v)
	}
	return ret
}

//隐藏中间几位
func StrHideCharacter(str string) string {
	if str == "" {
		return ""
	}
	strLen := len(str)
	strL3 := strLen / 3
	if strL3 > 4 {
		strL3 = 4
	}
	if strLen < 3 {
		return str
	}

	return str[0:strL3] + "****" + str[strLen-strL3:strLen]
}

//隐藏后面几位
func StrHideCharacterLast(str string, maxLen int) string {
	if str == "" {
		return ""
	}
	strLen := len(str)
	if strLen < 2 {
		return str
	}

	strL3 := strLen / 2
	if strL3 > maxLen {
		return str[0:maxLen] + "****"
	}

	return str[0:strL3] + "****"
}
