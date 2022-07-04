package payjava

import (
	"common-go/common/utils"
	"fmt"
	"github.com/go-resty/resty/v2"
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
