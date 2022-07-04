package redisVar

//redis key常量前缀
const (
	RedisTestLock = "BA:redisTestLock" //redis测试

	LoginMobileLock        = "BA:loginLock:mobile:"        //登陆手机号锁
	OrderNotifyOrderNoLock = "BA:orderNotifyLock:orderNo:" //支付通知锁
	OrderSharingCronLock   = "BA:OrderSharingCronLock:12"  //12点定时任务分账锁
)
