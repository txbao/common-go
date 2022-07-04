package authority

import (
	"bank-activity/common/utils"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

//会员登录
const (
	RedisKeyTokenPrefix = "BA:token:"
)

type TokenData struct {
	ActivityId int64  `json:"activity_id"` //活动ID
	AppId      int64  `json:"app_id"`      //提现后台用
	UserId     int64  `json:"user_id"`
	UserName   string `json:"user_name"`
	AccountId  int64  `json:"account_id"`
	WxOpenid   string `json:"wx_openid"` //发送立减金用
	CreateTime int64  `json:"create_time"`
	RequestNo  string `json:"request_no"` //汉口银行用
	Mobile     string `json:"mobile"`
}
type _token struct {
	comm _comm
}

func NewToken(rds *redis.Redis, expire int64) _token {
	return _token{
		comm: NewComm(rds, expire, RedisKeyTokenPrefix),
	}
}

// 生成token
func (obj *_token) TokenGenerate(data TokenData) (string, int64, error) {
	tokenJson, err := utils.StructToJsonStr(data)
	if err != nil {
		return "", 0, err
	}
	return obj.comm.TokenGenerate(tokenJson)
}

// token更新
func (obj *_token) TokenUpdate(token string, data TokenData) (string, int64, error) {
	tokenJson, err := utils.StructToJsonStr(data)
	if err != nil {
		return "", 0, err
	}
	return obj.comm.TokenUpdate(token, tokenJson)
}

// 获取token
func (obj *_token) GetToken(token string) (*TokenData, error) {
	tokenJson, err := obj.comm.GetToken(token)
	if err != nil {
		return nil, err
	}
	var data TokenData
	err = json.Unmarshal([]byte(tokenJson), &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

/*
// 检验token
func (obj *_token) TokenValid(token string, bankId int) error {
	tokenStruct, err := obj.GetToken(token)
	if err != nil {
		return err
	}
	if bankId != tokenStruct.BankId {
		return errorrpc.New("当前银行ID与Token不符")
	}
	return nil
}
*/

// 延期token
func (obj *_token) TokenExp(token string) error {
	return obj.comm.TokenExp(token)
}

//获取token的数据
func (obj *_token) GetTokenData(Authorization string) (*TokenData, error) {
	tokenJson, err := obj.comm.GetTokenData(Authorization)
	if err != nil {
		return nil, err
	}

	var data TokenData
	err = json.Unmarshal([]byte(tokenJson), &data)
	if err != nil {
		return nil, err
	}
	return &data, nil

}
