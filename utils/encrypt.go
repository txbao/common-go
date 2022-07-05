/**
 * 加密
 * Create by whimp(whimp@189.cn)
 * Date: 2020\5\8 14:51
 */
package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
)

func Md5String(data string) string {
	md5 := md5.New()
	md5.Write([]byte(data))
	md5Data := md5.Sum([]byte(nil))
	return hex.EncodeToString(md5Data)
}

//func Md5(data []byte) string {
//	md5 := md5.New()
//	md5.Write(data)
//	md5Data := md5.Sum([]byte(nil))
//	return hex.EncodeToString(md5Data)
//}

func HmacString(key string, data string) string {
	hmac := hmac.New(md5.New, []byte(key))
	hmac.Write([]byte(data))
	return hex.EncodeToString(hmac.Sum([]byte(nil)))
}

func HmacByStringKey(key string, data []byte) string {
	return Hmac([]byte(key), data)
}

func Hmac(key []byte, data []byte) string {
	hmac := hmac.New(md5.New, key)
	hmac.Write(data)
	return hex.EncodeToString(hmac.Sum([]byte(nil)))
}

func Sha1String(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte(nil)))
}

func Sha1(data []byte) string {
	sha1 := sha1.New()
	sha1.Write(data)
	return hex.EncodeToString(sha1.Sum([]byte(nil)))

}

// hmacsha256验证
func HMAC_SHA256(src, key string) string {
	m := hmac.New(sha256.New, []byte(key))
	m.Write([]byte(src))
	return hex.EncodeToString(m.Sum(nil))
}

// hmacsha512验证
func HMAC_SHA512(src, key string) string {
	m := hmac.New(sha512.New, []byte(key))
	m.Write([]byte(src))
	return hex.EncodeToString(m.Sum(nil))
}

// sha256验证
func SHA256Str(src string) string {
	h := sha256.New()
	h.Write([]byte(src)) // 需要加密的字符串为
	// fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果
	return hex.EncodeToString(h.Sum(nil))
}

// sha512验证
func SHA512Str(src string) string {
	h := sha512.New()
	h.Write([]byte(src)) // 需要加密的字符串为
	// fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果
	return hex.EncodeToString(h.Sum(nil))
}

// base编码
func BASE64EncodeStr(src string) string {
	return string(base64.StdEncoding.EncodeToString([]byte(src)))
}

// base解码
func BASE64DecodeStr(src string) string {
	a, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return ""
	}
	return string(a)
}


func CreateSign(params map[string]interface{}, appSecret string) string {
	var key []string
	var str = ""
	for k := range params {
		if k != "sign" {
			key = append(key, k)
		}
	}
	key = append(key, "app_secret")
	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		if key[i] == "app_secret" {
			str = str + appSecret
		} else {
			switch f := params[key[i]].(type) {
			case float64:
				str = str + Float64ToString(f)
			case map[string]interface{}:
				paramsJson, _ := json.Marshal(f)
				str = str + string(paramsJson)
			default:
				str = str + fmt.Sprintf("%v", f)
			}
		}
	}
	//判断时间戳是否超过5分钟
	//if params["timestamp"] != nil {
	//	timeStamp, err := utils.StringToInt64(params["timestamp"].(string))
	//	if err != nil {
	//		return ""
	//	}
	//	nowTime := time.Now().UnixNano() / 1e6
	//	if nowTime-timeStamp > 300000 {
	//		return "timeout"
	//	}
	//}
	// 自定义签名算法
	fmt.Println("加密字符串===",str)
	sign := Md5String(str)
	return sign
}