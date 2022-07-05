package payment

import (
	"bytes"
	"common-go/common/utils"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

//开户调试日志
const PAY_DEBUG = false

// 加密
func (p *PayAppParams) Encrypt(param url.Values) string {
	payPrivateKey := p.PrivateKey // viper.GetString("settings.payment.privateKey")
	payPublicKey := p.PublicKey

	if PAY_DEBUG {
		fmt.Println("--------------------")
		fmt.Println("payPrivateKey", payPrivateKey)
		fmt.Println("--------------------")
		fmt.Println("payPublicKey", payPublicKey)
	}

	privateKey := bytes.NewBufferString(utils.FormatPrivateKey(payPrivateKey))
	publicKey := bytes.NewBufferString(utils.FormatPublicKey(payPublicKey))
	xrsa, err := utils.XrsaNewXRsa(publicKey.Bytes(), privateKey.Bytes())
	if err != nil {
		fmt.Println("--------------------")
		fmt.Println("err", err)
		return ""
	}

	var pList = make([]string, 0, 0)
	for key := range param {
		var value = strings.TrimSpace(param.Get(key))
		if len(value) > 0 && key != "sign" {
			//去除空格，因为空格与+号，加密是会报错
			value = StrReplaceEncrypt(value)
			key = StrReplaceEncrypt(key)

			pList = append(pList, key+"="+value)
		}
	}
	sort.Strings(pList)
	var signStr = strings.Join(pList, "&")
	if PAY_DEBUG {
		fmt.Println("--------------------")
		fmt.Println("加密前signStr：", signStr)
	}
	sign, _ := xrsa.XrsaSign(signStr)
	param.Add("sign", sign)
	signStr += "&sign=" + sign
	data, _ := xrsa.XrsaPublicEncrypt(signStr)

	if PAY_DEBUG {
		fmt.Println("--------------------")
		fmt.Println("加密后signStr：", signStr)
	}

	return data
}

//解密
func (p *PayAppParams) Decrypt(data string) (url.Values, error) {
	payPrivateKey := p.PrivateKey // viper.GetString("settings.payment.privateKey")
	payPublicKey := p.PublicKey
	if PAY_DEBUG {
		fmt.Println("--------------------")
		fmt.Println("payPrivateKey", payPrivateKey)
		fmt.Println("--------------------")
		fmt.Println("payPublicKey", payPublicKey)
	}

	privateKey := bytes.NewBufferString(utils.FormatPrivateKey(payPrivateKey))
	publicKey := bytes.NewBufferString(utils.FormatPublicKey(payPublicKey))
	xrsa, err := utils.XrsaNewXRsa(publicKey.Bytes(), privateKey.Bytes())
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	decrypted, _ := xrsa.XrsaPrivateDecrypt(data)

	decryptedMap, _ := url.ParseQuery(decrypted)
	sign := decryptedMap["sign"][0]

	var pList = make([]string, 0, 0)
	for key := range decryptedMap {
		var value = strings.TrimSpace(decryptedMap.Get(key))
		if len(value) > 0 && key != "sign" {
			pList = append(pList, key+"="+value)
		}
	}
	sort.Strings(pList)
	var signStr = strings.Join(pList, "&")
	err = xrsa.XrsaVerify(signStr, sign)

	if PAY_DEBUG {
		fmt.Println("--------------------")
		fmt.Println("decrypted", decrypted, "")
		fmt.Println("--------------------")
		fmt.Println("sign", sign)
		fmt.Println("--------------------")
		fmt.Println("signStr", signStr)
		fmt.Println("--------------------")
		fmt.Println("解密", decryptedMap)
	}

	return decryptedMap, err
}
