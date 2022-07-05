package captchas

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

const CAPTCHA = "captcha:"

type RedisStore struct {
	rds *redis.Redis
}

// 存储的是 "8926" 而不是 8926，需要转换一下
func (r RedisStore) Set(id string, value string) error {
	//v ,_ := strconv.Atoi(value)
	key := CAPTCHA + ":" + id
	err := r.rds.Setex(key, value, 180)
	if err != nil {
		fmt.Println("验证码保存失败:", err.Error())
		return err
	}
	//databases.Redis.Set(CAPTCHA + global.Config.Common.Name + ":" + id, value, 180 * time.Second)
	return nil
}

func (r RedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + ":" + id
	val, err := r.rds.Get(key)
	if err != nil {
		return ""
	}
	if clear {
		_, err := r.rds.Del(key)
		if err != nil {
			fmt.Println("验证码删除失败:", err.Error())
		}
	}

	return val
}

func (r RedisStore) Verify(id string, answer string, clear bool) bool {
	//v := RedisStore{}.Get(id, clear)
	v := r.Get(id, clear)
	return v == answer
}
