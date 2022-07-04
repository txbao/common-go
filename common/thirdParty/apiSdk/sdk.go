package apiSdk

import (
	"common-go/common/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

//对外接口API

type _sdk struct {
	xrsa *utils.XRsa
}

type ReqStruct struct {
	AppId int64  `json:"app_id"` //应用ID
	Data  string `json:"data"`   //接口参数（json格式）加密后字符串
	Sign  string `json:"sign"`   //签名
}

type ReqDataStruct struct {
	AppId     int64       `json:"app_id"`    //应用ID
	Service   string      `json:"service"`   //服务
	Method    string      `json:"method"`    //方法
	Timestamp int64       `json:"timestamp"` //时间戳，如1569135887
	Data      interface{} `json:"data"`
}

//通用格式化
func formatKey(key string) string {
	keyLen := len(key)
	keyStr := "\n"
	for i := 0; i <= keyLen; i = i + 64 {
		s := 0 + i
		e := 64 + i
		if e > keyLen {
			e = keyLen
		}
		keyStr += key[s:e] + "\n"
	}
	return keyStr
}

//格式化私钥
func FormatPrivateKey(privateKey string) string {
	return "-----BEGIN RSA PRIVATE KEY-----" + formatKey(privateKey) + "-----END RSA PRIVATE KEY-----"
}

// 格式化公钥
func FormatPublicKey(publicKey string) string {
	return "-----BEGIN PUBLIC KEY-----" + formatKey(publicKey) + "-----END PUBLIC KEY-----"
}

// map转 json str
func Map2Json(result interface{}) string {
	jsonBytes, _ := json.Marshal(result)
	jsonStr := string(jsonBytes)
	return jsonStr
}
func StructToJsonStr(e interface{}) (string, error) {
	if b, err := json.Marshal(e); err == nil {
		return string(b), err
	} else {
		return "", err
	}
}

func NewSdk(appId int64, privateKey string, publicKey string) (*_sdk, error) {
	priKey := bytes.NewBufferString(FormatPrivateKey(privateKey))
	pubKey := bytes.NewBufferString(FormatPublicKey(publicKey))

	var err error
	xrsa, err := utils.XrsaNewXRsa(pubKey.Bytes(), priKey.Bytes())
	if err != nil {
		return nil, err
	}
	return &_sdk{
		xrsa: xrsa,
	}, nil
}

//post请求
func (obj *_sdk) HttpPost(apiUrl string, params map[string]interface{}) (string, error) {
	client := resty.New()
	resp, err := client.R().
		SetBody(params).
		SetHeader("Content-Type", "application/json").
		Post(apiUrl)
	if err != nil {
		return "", err
	}
	return resp.String(), err
}

//请求
func (obj *_sdk) Request(reqUrl string, appId int64, reqJson string) (string, error) {
	sign, _ := obj.xrsa.XrsaSign(reqJson)
	data, _ := obj.xrsa.XrsaPublicEncrypt(reqJson)

	//组织请求参数
	reqMp := make(map[string]interface{})
	reqMp["app_id"] = appId
	reqMp["data"] = data
	reqMp["sign"] = sign
	res, err := obj.HttpPost(reqUrl, reqMp)
	fmt.Println("请求结果:", res, "请求地址：", reqUrl, "请求数据：", utils.Map2Json(reqMp), "加密数据：", reqJson)

	return res, err
}

//RSA验证
func (obj *_sdk) VerifyRsa(Data string, sign string) (string, error) {
	decrypted, err := obj.xrsa.XrsaPrivateDecrypt(Data)
	if err != nil {
		return "", err
	}

	err = obj.xrsa.XrsaVerify(decrypted, sign)
	if err != nil {
		return "", err
	}
	return decrypted, nil
}
