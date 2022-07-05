package aliyunVenuscn

import (
	"github.com/txbao/common-go/common/utils"
	"github.com/txbao/common-go/common/utils/logs"
	"fmt"
	"sync"
	"time"
)

//阿里云-手机号码归属地查询-深圳华辰网络科技有限公司
const (
	Path   = "/mobile"
	Method = "GET"
)

type ResStruct struct {
	Data struct {
		AreaCode string `json:"area_code"`
		City     string `json:"city"`
		CityCode string `json:"city_code"`
		Isp      string `json:"isp"`
		Lat      string `json:"lat"`
		Lng      string `json:"lng"`
		Num      int32  `json:"num"`
		Prov     string `json:"prov"`
		Types    string `json:"types"`
		ZipCode  string `json:"zip_code"`
	} `json:"data"`
	LogID string `json:"log_id"`
	Msg   string `json:"msg"`
	Ret   int64  `json:"ret"`
}

type Config struct {
	once    sync.Once
	host    string
	appcode string
}

func NewSdk(host string, appcode string) *Config {
	return &Config{
		host:    host,
		appcode: appcode,
	}

}

//请求
func (c *Config) Request(args map[string]interface{}) (ResStruct, error) {
	//记录开始时间
	start := time.Now() // 获取当前时间

	headerMap := make(map[string]interface{})
	headerMap["Authorization"] = fmt.Sprintf("APPCODE %s", o.appcode)

	var res ResStruct
	resStr, err := utils.HttpGetResty(o.host+Path, args, headerMap, &res)
	//记录结束时间
	elapsed := time.Since(start)
	if err != nil {
		logs.Info("阿里云-手机号归属请求数据：req:", utils.Map2Json(args), "\n", "响应数据：", err.Error(), "\n执行完成耗时：", elapsed)
		return res, err
	}
	logs.Info("阿里云-手机号归属请求数据：req:", utils.Map2Json(args), "\n", "响应数据：", resStr.String(), "\n执行完成耗时：", elapsed)

	return res, err
}
