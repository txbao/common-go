package express

import (
	"common-go/utils"
	"common-go/utils/logs"
	"fmt"
	"time"
)

const (
	PathQuery = "/express/query"
	PathType  = "/express/type"
	Method    = "ANY"
)

type ExpressQuery struct {
	Msg    string `json:"msg"`
	Result struct {
		Deliverystatus int64 `json:"deliverystatus"`
		Issign         int64 `json:"issign"`
		List           []struct {
			Status string `json:"status"`
			Time   string `json:"time"`
		} `json:"list"`
		Logo     string `json:"logo"`
		Number   string `json:"number"`
		Type     string `json:"type"`
		Typename string `json:"typename"`
	} `json:"result"`
	Status int64 `json:"status"`
}

type ExpressType struct {
	Msg    string `json:"msg"`
	Result []struct {
		Letter string `json:"letter"`
		Name   string `json:"name"`
		Number string `json:"number"`
		Tel    string `json:"tel"`
		Type   string `json:"type"`
		Logo   string `json:"logo"`
	} `json:"result"`
	Status int64 `json:"status"`
}

type expressSdk struct {
	appCode string
	host    string
}

func NewSdk(appCode string, host string) (*expressSdk, error) {
	return &expressSdk{
		appCode: appCode,
		host:    host,
	}, nil
}

//获取快递类型
func (o *expressSdk) GetPressType() (*ExpressType, error) {
	args := make(map[string]interface{})
	var expressType ExpressType

	//记录开始时间
	start := time.Now() // 获取当前时间
	headerMap := make(map[string]interface{})
	headerMap["Authorization"] = fmt.Sprintf("APPCODE %s", o.appCode)
	resStr, err := utils.HttpGetResty(o.host+PathType, args, headerMap, &expressType)
	//记录结束时间
	elapsed := time.Since(start)
	if err != nil {
		logs.Info("阿里云-物流类型查询请求数据：req:", utils.Map2Json(args), "\n", "响应数据：", err.Error(), "\n执行完成耗时：", elapsed)
		return nil, err
	}

	logs.Info("阿里云-物流类型查询请求数据：req:", utils.Map2Json(args), "\n", "响应数据：", resStr.String(), "\n执行完成耗时：", elapsed)
	return &expressType, nil
}

//获取快递类型
func (o *expressSdk) GetPressQuery(number string, types string, mobile string) (string, error) {
	number = utils.TrimStr(number)
	types = utils.TrimStr(types)

	if types == "" {
		types = "auto"
	}
	//因为部分快递容易选择错，目前用自动识别方式
	types = "auto"
	args := make(map[string]interface{})
	args["mobile"] = utils.StringGetLastChar(mobile, 4) //可选	收件人/寄件人手机号（顺丰快递需要） 后四位
	args["number"] = number                             //必选	快递单号
	args["type"] = types                                //必选	快递公司 自动识别请写auto

	//记录开始时间
	start := time.Now() // 获取当前时间
	headerMap := make(map[string]interface{})
	headerMap["Authorization"] = fmt.Sprintf("APPCODE %s", o.appCode)
	resStr, err := utils.HttpGetResty(Host+PathQuery, args, headerMap, nil)
	//记录结束时间
	elapsed := time.Since(start)
	if err != nil {
		logs.Info("阿里云-物流查询请求数据：req:", utils.Map2Json(args), "\n", "响应数据：", err.Error(), "\n执行完成耗时：", elapsed)
		return "", err
	}
	logs.Info("阿里云-物流查询请求数据：req:", utils.Map2Json(args), "\n", "响应数据：", resStr.String(), "\n执行完成耗时：", elapsed)
	return resStr.String(), nil
}
