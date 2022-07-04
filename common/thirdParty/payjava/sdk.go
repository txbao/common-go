package payjava

import (
	"bank-activity/common/utils"
	"fmt"
	"github.com/go-resty/resty/v2"
)

const (
	//中行解密URL
	BocDecryptMessageUrl = "/tools/boc/decryptmessage"
	//中行解密公钥私钥
	BocPriKey     = "MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQgfmFP0BtuOWhbZ1eSRzvSxinuUVxEeY6VXNV8h9LCqhSgCgYIKoEcz1UBgi2hRANCAASEh545ndeZbFnrQHvZW+w8rOsZ+yKQr5+1KsSgiPe0lVqVbduyIP/YHABA8c5fAAi11B5heDSoMjDS1eY8/+dn"
	BocPubKey     = "MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEhIeeOZ3XmWxZ60B72VvsPKzrGfsikK+ftSrEoIj3tJValW3bsiD/2BwAQPHOXwAItdQeYXg0qDIw0tXmPP/nZw=="
	PubKeyForHmac = "MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEqcZ4DtUvlRpSvwvhZOIkQsSmuIT3FMmOi5dR08tEER85yrZ9uXRaR6f7v1VePKl9d2lkUqnXLWU9dgpvg705Ng=="
)

type PayJava struct {
	TimeStamp  int64
	RequestUrl string
	AppKey     string
	AppSecret  string
}

func NewSdk(timeStamp int64, requestUrl string, appKey string, appSecret string) *PayJava {
	return &PayJava{
		TimeStamp:  timeStamp,
		RequestUrl: requestUrl,
		AppKey:     appKey,
		AppSecret:  appSecret,
	}
}

//加密
func (obj *PayJava) sign() string {
	signStr := fmt.Sprintf("%s&%s&%v", obj.AppKey, obj.AppSecret, obj.TimeStamp)
	fmt.Println("signStr", signStr)
	return utils.Md5(signStr)
}

//header
func (obj *PayJava) header() map[string]interface{} {
	fmt.Println("obj.AppKey", obj.AppKey)
	headers := map[string]interface{}{
		"sign":      obj.sign(),
		"appkey":    obj.AppKey,
		"timestamp": utils.Int64ToString(obj.TimeStamp),
	}
	return headers
}

//请求
func (obj *PayJava) Request(gateway string, params map[string]interface{}, result interface{}) (*resty.Response, error) {
	requestUrl := obj.RequestUrl + gateway
	fmt.Println("requestUrl", requestUrl)
	headers := obj.header()
	return utils.HttpPostResty(requestUrl, params, headers, result)
}
