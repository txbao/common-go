package aliyunVenuscn

import (
	"bank-activity/common/utils"
	"bank-activity/common/utils/logs"
	"fmt"
	"sync"
	"time"
)

//阿里云-手机号码归属地查询-深圳华辰网络科技有限公司

//AppKey：203753318
//AppSecret：y4880p7dn55sy8lbyvhecjqo3u78winf
//AppCode：824b02b71fec46078ebaee788753e7cc
const (
	Host    = "https://api04.aliyun.venuscn.com"
	Path    = "/mobile"
	Method  = "GET"
	Appcode = "824b02b71fec46078ebaee788753e7cc"
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
	once sync.Once
}

func NewSdk() *Config {
	return &Config{}

}

//请求
func (c *Config) Request(args map[string]interface{}) (ResStruct, error) {
	//记录开始时间
	start := time.Now() // 获取当前时间

	headerMap := make(map[string]interface{})
	headerMap["Authorization"] = fmt.Sprintf("APPCODE %s", Appcode)

	var res ResStruct
	resStr, err := utils.HttpGetResty(Host+Path, args, headerMap, &res)
	//记录结束时间
	elapsed := time.Since(start)
	if err != nil {
		logs.Info("阿里云-手机号归属请求数据：req:", utils.Map2Json(args), "\n", "响应数据：", err.Error(), "\n执行完成耗时：", elapsed)
		return res, err
	}
	logs.Info("阿里云-手机号归属请求数据：req:", utils.Map2Json(args), "\n", "响应数据：", resStr.String(), "\n执行完成耗时：", elapsed)

	return res, err
}
