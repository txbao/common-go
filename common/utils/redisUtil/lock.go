package redisUtil

import (
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

//Redis锁
func RedisLock(rds *redis.Redis, redisLockKey string, expire int) error {
	if expire == 0 {
		expire = 5
	}
	// 1. New redislock
	redisLock := redis.NewRedisLock(rds, redisLockKey)
	// 2. 可选操作，设置 redislock 过期时间
	redisLock.SetExpire(expire)
	if ok, err := redisLock.Acquire(); !ok || err != nil {
		return errors.New("请求频繁，请稍后再试！")
	}
	//defer func() {
	//	recover()
	//	// 3. 释放锁
	//	redisLock.Release()
	//}()

	return nil
}
