package code

import (
	"errors"
	"fmt"
	"github.com/txbao/common-go/utils"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

//兑换码

const (
	PreDigit = 3 //前缀位数 26*26*26-1
	CharStr  = "123456789ABCDEFGHJKLMNPQRSTUVWXYZ"
)

type _code struct {
}

func NewSdk() *_code {
	return &_code{}
}

//前缀编码
var chars string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func encode(num int64) string {
	bytes := []byte{}
	for num > 0 {
		bytes = append(bytes, chars[num%36])
		num = num / 36
	}
	reverse(bytes)
	return string(bytes)
}
func decode(str string) int64 {
	var num int64
	n := len(str)
	for i := 0; i < n; i++ {
		pos := strings.IndexByte(chars, str[i])
		num += int64(math.Pow(36, float64(n-i-1)) * float64(pos))
	}
	return num
}
func reverse(a []byte) {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
}

//获取最大可生成券码数
func GetMax(digit int) int64 {
	var num int = len(CharStr)
	var ji int = 1
	for i := 0; i < digit; i++ {
		fmt.Println("A:", i)
		ji *= num
	}
	return int64(ji)
}

//获取前缀
func (o *_code) getCodePre(id int64) (string, error) {
	if id > 46655 {
		return "", errors.New("最多只能生成46655次券码")
	}
	pre := encode(id)
	for i := len(pre); i < PreDigit; i++ {
		pre = "0" + pre
	}
	return pre, nil
}

//获取前缀
func (o *_code) GetCodePre(id int64) (string, error) {
	if id > 46655 {
		return "", errors.New("最多只能生成46655次券码")
	}
	pre := encode(id)
	for i := len(pre); i < PreDigit; i++ {
		pre = "0" + pre
	}
	return pre, nil
}

//获取字符串的Ascii值,区分数字与字母
func (o *_code) getOrd(str string) int64 {
	var sum int64 = 0
	for _, v := range str {
		if unicode.IsDigit(v) {
			v64, _ := strconv.ParseInt(string(v), 10, 64)
			sum += v64
		} else {
			sum += int64(utils.Ord(string(v)))
		}
	}
	return sum
}

//获取追加字符
func (o *_code) GetAdditionalChar(i int64) (string, error) {
	if i < 0 {
		return "", errors.New("追加数不能小于0")
	}
	//str := "1,2,3,4,5,6,7,8,9,A,B,C,D,E,F,G,H,J,K,L,M,N,P,Q,R,S,T,U,V,W,X,Y,Z"
	str := ""
	for i := 0; i < len(CharStr); i++ {
		if i != 0 {
			str = fmt.Sprintf("%s%s", str, ",")
		}
		str = fmt.Sprintf("%s%s", str, CharStr[i:i+1])
	}

	strArr := strings.Split(str, ",")
	if int(i) >= len(strArr) {
		return "", errors.New("最大只能追加33个")
	}
	return strArr[i], nil
}

//生成兑换码
func (o *_code) Create(total int, digit int, batchId int64, preStr string, additionalTimes int64) ([]string, error) {
	//获取追加字符串
	additional, err := o.GetAdditionalChar(additionalTimes)
	if err != nil {
		return nil, err
	}
	//减去前缀
	digit = digit - 2        //预留首位+末位加密码
	digit = digit - PreDigit //预留3位顺序位
	digit = digit - 1        //预留一位追加券
	pre, err := o.getCodePre(batchId)
	if err != nil {
		return nil, err
	}

	var codeArr []string
	//var b58Alphabet = []byte("123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")
	var b58Alphabet = []byte(CharStr)
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
		fmt.Println("idstr=", idstr)
		// 去除重复
		if _, ok := m[idstr]; ok {
			continue
		}
		m[idstr] = struct{}{}

		idstrEncrypt := preStr + o.encrypt(idstr+additional, pre)
		codeArr = append(codeArr, idstrEncrypt)

		a++
		if a >= total {
			break
		}
	}
	result := len(m)
	fmt.Println("生成", total)
	fmt.Printf("重复率%f\n", float64(total-result)/float64(total))
	fmt.Println("用时", time.Since(t))
	return codeArr, nil
}

//加密方法
func (o *_code) encrypt(str string, pre string) string {
	var sumCode string = str
	var sum int64 = o.getOrd(sumCode)
	//fmt.Println("sum:", sum)
	var preInt int64 = o.getOrd(pre)
	sum = (sum + preInt) * 3
	sum = sum ^ 28
	sum16 := utils.IntHex10to16(int(sum + preInt*5))
	strStartChar := sum16[0:1]
	sumStr := utils.Int64ToString(sum)
	sumStrLen := len(sumStr)

	if utils.IsNumeric(pre+str) && !(utils.IsNumeric(strStartChar)) {
		strStartChar = sumStr[0:1]
	}
	strEndChar := sumStr[sumStrLen-1 : sumStrLen]

	fmt.Printf("strStartChar:%v , pre:%v , str：%v , strEndChar:%v", strStartChar, pre, str, strEndChar)
	code := strStartChar + pre + str + strEndChar
	return code
}

//验证
func (o *_code) Verify(code string) bool {
	if code == "" {
		return false
	}
	codeLen := len(code)

	strStartCharV := code[0:1]
	strEndCharV := code[codeLen-1 : codeLen]
	pre := code[1 : PreDigit+1]

	var sumCode string = code[PreDigit+1 : codeLen-1]
	var sum int64 = o.getOrd(sumCode)
	//fmt.Println("sum:", sum)
	var preInt int64 = o.getOrd(pre)

	sum = (sum + preInt) * 3
	sum = sum ^ 28
	sum_16 := utils.IntHex10to16(int(sum + preInt*5))
	strStartChar := sum_16[0:1]
	sumStr := utils.Int64ToString(sum)
	sumStrLen := len(sumStr)

	if utils.IsNumeric(code) && !(utils.IsNumeric(strStartChar)) {
		strStartChar = sumStr[0:1]
	}

	strEndChar := sumStr[sumStrLen-1 : sumStrLen]
	//fmt.Println("校验：",strings.ToUpper(strStartChar),strings.ToUpper(strStartCharV),strings.ToUpper(strEndCharV),strings.ToUpper(strEndChar),"\n")
	if strings.ToUpper(strStartChar) == strings.ToUpper(strStartCharV) && strings.ToUpper(strEndCharV) == strings.ToUpper(strEndChar) {
		return true
	}
	return false
}
