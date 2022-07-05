package authority

import (
	"common-go/utils"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"time"
)

const (
	AuthorizationTag = "Authorization"
)

type _comm struct {
	rds                 *redis.Redis
	expire              int64
	redisKeyTokenPrefix string
}

func NewComm(rds *redis.Redis, expire int64, redisKeyTokenPrefix string) _comm {
	return _comm{
		rds:                 rds,
		expire:              expire,
		redisKeyTokenPrefix: redisKeyTokenPrefix,
	}
}

// 生成token
func (obj *_comm) TokenGenerate(tokenJson string) (string, int64, error) {
	mm, _ := time.ParseDuration(utils.Int64ToString(obj.expire))
	mm1 := time.Now().Add(mm)
	expire := mm1.Unix()

	token := uuid.New().String()
	err := obj.rds.Setex(obj.redisKeyTokenPrefix+token, tokenJson, int(obj.expire))
	if err != nil {
		return "", 0, err
	}
	return token, expire, nil
}

// token更新
func (obj *_comm) TokenUpdate(token string, tokenJson string) (string, int64, error) {
	mm, _ := time.ParseDuration(utils.Int64ToString(obj.expire))
	mm1 := time.Now().Add(mm)
	expire := mm1.Unix()

	err := obj.rds.Setex(obj.redisKeyTokenPrefix+token, tokenJson, int(obj.expire))
	if err != nil {
		return "", 0, err
	}
	return token, expire, nil
}

// 获取token
func (obj *_comm) GetToken(token string) (string, error) {
	tokenJson, err := obj.rds.Get(obj.redisKeyTokenPrefix + token)
	if err != nil {
		return "", err
	}
	if tokenJson == "" {
		//return nil, errors.New("Token is expired")
		return "", errors.New("登录超时，请重新登录！")
	}
	fmt.Println("tokenJson", tokenJson)
	return tokenJson, nil
}

/*
// 检验token
func (obj *_comm) TokenValid(token string, bankId int) error {
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
func (obj *_comm) TokenExp(token string) error {
	err := obj.rds.Expire(obj.redisKeyTokenPrefix+token, int(obj.expire))
	return err
}

//获取token的数据
func (obj *_comm) GetTokenData(Authorization string) (string, error) {
	if Authorization == "" {
		return "", errors.New("Authorization不能为空")
	}
	return obj.GetToken(Authorization)
}
