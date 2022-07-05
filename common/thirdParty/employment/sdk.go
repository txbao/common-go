package employment

//灵活用工

import (
	"bytes"
	"common-go/common/utils"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

type Epy struct {
	TimeStamp int64
}

//查询签约状态结构体
type GetSignStatusStruct struct {
	CompanyCode string `json:"company_code"`
	IdentityNo  string `json:"identity_no"`
	TaskCode    string `json:"task_code"`
	Sign        string `json:"sign"`
}

//提现申请结构体
type ApplyIssuingStruct struct {
	CompanyCode string `json:"company_code"`
	IdentityNo  string `json:"identity_no"`
	TaskCode    string `json:"task_code"`
	OrderCode   string `json:"order_code"`
	Count       int    `json:"count"`
	Amount      string `json:"amount"`
	Location    string `json:"location"`
	Content     string `json:"content"`
	Name        string `json:"name"`
	SubmitTime  string `json:"submit_time"`
	Remark      string `json:"remark"`
	PhotoUrl    string `json:"photo_url"`
	Sign        string `json:"sign"`
}

//请求结构
type ReqStruct struct {
	Gateway string      `json:"gateway"`
	Url     string      `json:"url"`
	Data    interface{} `json:"data"`
}

//返回结构
type ResStruct struct {
	Code int `json:"code"`
	Data struct {
		SignStatus int    `json:"sign_status"`
		Mobile     string `json:"mobile"`
		Name       string `json:"name"`
		IdentityNo string `json:"identity_no"`
		BankNo     string `json:"bank_no"`
	} `json:"data"`
	Msg          string `json:"msg"`
	BusinessCode int64  `json:"statusCode"`
	Success      bool   `json:"success"`
}

//提现结果返回结构
type GetApplyResStruct struct {
	Code int `json:"code"`
	Data struct {
		PayTime   string `json:"pay_time"`
		Amount    string `json:"amount"`
		Remark    string `json:"remark"`
		PayStatus int    `json:"pay_status"`
	} `json:"data"`
	Msg          string `json:"msg"`
	BusinessCode int64  `json:"statusCode"`
	Success      bool   `json:"success"`
}

//修改银行卡信息返回结构
type EditBankcardResStruct struct {
	Code int `json:"code"`
	Data struct {
	} `json:"data"`
	Msg          string `json:"msg"`
	BusinessCode int64  `json:"statusCode"`
	Success      bool   `json:"success"`
}

//提现返回结构体
type WithdrawResStruct struct {
	BusinessCode int64              `json:"businessCode"`
	Code         int64              `json:"code"`
	Data         WithdrawDataStruct `json:"data"`
	Msg          string             `json:"msg"`
	Success      bool               `json:"success"`
}
type WithdrawDataStruct struct {
}

type _sdk struct {
	Xrsa *utils.XRsa
}

//创建实例结构体
func NewSdk(privateKey string, publicKey string) (*_sdk, error) {
	var err error
	priKey := bytes.NewBufferString(utils.FormatPrivateKey(privateKey))
	pubKey := bytes.NewBufferString(utils.FormatPublicKey(publicKey))

	XRsa, err := utils.XrsaNewXRsa(pubKey.Bytes(), priKey.Bytes())
	if err != nil {
		fmt.Println("employment 初始化 err::", err)
		return nil, err
	}
	return &_sdk{
		Xrsa: XRsa,
	}, nil
}

//开户调试日志
const PAY_DEBUG = true

// 加密
func (o *_sdk) EmpEncrypt(param url.Values) string {
	var pList = make([]string, 0, 0)
	for key := range param {
		var value = strings.TrimSpace(param.Get(key))
		if len(value) > 0 && key != "sign" {
			//去除空格，因为空格与+号，加密是会报错
			//value = StrReplaceEncrypt(value)
			//key = StrReplaceEncrypt(key)

			pList = append(pList, key+"="+value)
		}
	}
	sort.Strings(pList)
	var signStr = strings.Join(pList, "&")
	if PAY_DEBUG {
		fmt.Println("--------------------")
		fmt.Println("加密前signStr：", signStr)
	}
	sign, _ := o.Xrsa.XrsaSign(signStr)
	param.Add("sign", sign)
	signStr += "&sign=" + sign
	data, _ := o.Xrsa.XrsaPublicEncrypt(signStr)

	if PAY_DEBUG {
		fmt.Println("--------------------")
		fmt.Println("加密后signStr：", signStr)
	}

	return data
}

//解密
func (o *_sdk) EmpDecrypt(data string) (url.Values, error) {
	decrypted, _ := o.Xrsa.XrsaPrivateDecrypt(data)

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
	err := o.Xrsa.XrsaVerify(signStr, sign)

	if PAY_DEBUG {
		fmt.Println("解密", decryptedMap)
	}

	return decryptedMap, err
}

//加密的特殊字符替换， 如"+",空格等
func (o *_sdk) StrReplaceEncrypt(subject string) string {
	subject = o.StrReplace(" ", "", subject, -1)
	subject = o.StrReplace("+", "＋", subject, -1)
	return subject
}
func (o *_sdk) StrReplace(search, replace, subject string, count int) string {
	return strings.Replace(subject, search, replace, count)
}

// 加密
func (o *_sdk) CheckSign(param url.Values) string {
	var pList = make([]string, 0, 0)
	for key := range param {
		var value = strings.TrimSpace(param.Get(key))
		if len(value) > 0 && key != "sign" {
			//去除空格，因为空格与+号，加密是会报错
			//value = StrReplaceEncrypt(value)
			//key = StrReplaceEncrypt(key)

			pList = append(pList, key+"="+value)
		}
	}
	sort.Strings(pList)
	var signStr = strings.Join(pList, "&")
	if PAY_DEBUG {
		fmt.Println("--------------------")
		fmt.Println("加密前signStr：", signStr)
	}
	sign, _ := o.Xrsa.XrsaSign(signStr)
	return sign
}

func (o *_sdk) EncryptPostData(postStr string) string {
	data, _ := o.Xrsa.XrsaPublicEncrypt(postStr)
	return data
}

//提现回调验签
func (o *_sdk) VerifyNotify(data string) (url.Values, error) {
	dataUrl, err := o.EmpDecrypt(data)
	if err != nil {
		return nil, err
	}
	sign := o.CheckSign(dataUrl)
	if dataUrl.Get("sign") != sign {
		return dataUrl, err
	}
	return dataUrl, nil
}
