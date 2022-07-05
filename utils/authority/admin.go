package authority

import (
	"common-go/utils"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type AdminTokenData struct {
	UserId    int64  `json:"user_id"`
	UserName  string `json:"user_name"`
	UserToken string `json:"user_token"`
}
type _admin struct {
	comm _comm
}

func NewAdmin(rds *redis.Redis, expire int64, redisKeyPrefix string) _admin {
	return _admin{
		comm: NewComm(rds, expire, redisKeyPrefix),
	}
}

// 生成token
func (obj *_admin) TokenGenerate(data AdminTokenData) (string, int64, error) {
	tokenJson, err := utils.StructToJsonStr(data)
	if err != nil {
		return "", 0, err
	}
	return obj.comm.TokenGenerate(tokenJson)
}

// token更新
func (obj *_admin) TokenUpdate(token string, data AdminTokenData) (string, int64, error) {
	tokenJson, err := utils.StructToJsonStr(data)
	if err != nil {
		return "", 0, err
	}
	return obj.comm.TokenUpdate(token, tokenJson)
}

// 获取token
func (obj *_admin) GetToken(token string) (*AdminTokenData, error) {
	tokenJson, err := obj.comm.GetToken(token)
	if err != nil {
		return nil, err
	}
	var data AdminTokenData
	err = json.Unmarshal([]byte(tokenJson), &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

/*
// 检验token
func (obj *_admin) TokenValid(token string, bankId int) error {
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
func (obj *_admin) TokenExp(token string) error {
	return obj.comm.TokenExp(token)
}

//获取token的数据
func (obj *_admin) GetTokenData(Authorization string) (*AdminTokenData, error) {
	tokenJson, err := obj.comm.GetTokenData(Authorization)
	if err != nil {
		return nil, err
	}

	var data AdminTokenData
	err = json.Unmarshal([]byte(tokenJson), &data)
	if err != nil {
		return nil, err
	}
	return &data, nil

}
