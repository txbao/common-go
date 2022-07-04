package bocYn

import (
	"bank-activity/common/utils"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

//云南中行白名单接口
const (
	RedisKeyBocSearchCity = "bank:interface:bocyn:%s"
)

var ApiBocYN = &_boc{
	key: "ynboc1002",
}

type _boc struct {
	rds *redis.Redis
	key string
}

type BocSearchCityRes struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Res  string `json:"res"`
}

func NewSdk(rds *redis.Redis) *_boc {
	return &_boc{
		rds: rds,
		key: "ynboc1002",
	}
}

//白名单接口
func (o *_boc) SearchCity(phoneNum string) string {
	apiUrl := "https://cloud.bankofchina.com/yn/cardmobilecheck/mobile/searchcity"
	resStr, err := utils.HttpGet(fmt.Sprintf("%s?key=%s&phnno=%s", apiUrl, o.key, phoneNum))
	if err != nil {
		logx.Error("ApiBocYNSearchCityHttpGetError", map[string]interface{}{
			"err":      err,
			"phoneNum": phoneNum,
			"resStr":   resStr,
		})
		return ""
	}
	res := &BocSearchCityRes{}
	err = json.Unmarshal([]byte(resStr), res)
	if err != nil {
		logx.Error("ApiBocYNSearchCityUnmarshalError", map[string]interface{}{
			"err":      err,
			"phoneNum": phoneNum,
			"resStr":   resStr,
		})
		return ""
	}
	if res.Res == "否" {
		logx.Error("ApiBocYNSearchCityFalse", map[string]interface{}{
			"phoneNum": phoneNum,
			"resStr":   resStr,
		})
	}

	return res.Res
}

//通过缓存
func (o *_boc) SearchCityFromCache(phoneNum string) string {
	var res string
	redisKey := fmt.Sprintf(RedisKeyBocSearchCity, phoneNum)
	res, err := o.rds.Get(redisKey)
	if err != nil || res == "" {
		res = o.SearchCity(phoneNum)
		if res != "" {
			err = o.rds.Setex(redisKey, res, 3600)
			if err != nil {
				logx.Error("云南中行白名单保存Redis错误_BocYnErr：", err.Error())
			}
		}
	}

	if res == "否" {
		return ""
	} else {
		return res
	}
}
