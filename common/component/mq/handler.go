package mq

import "bank-activity/common/component/mq/rabbitmq"

// MQService mq 服务,目前暂定 kafka, 包含 send 和 read 两个方法
type MQService struct {
	// kafka.Kafka
	rabbitmq.Rabbitmq
}

// MQ 队列应该实现的方法
type MQ interface {
	Consum(func(jsonStr []byte))
	Publish(key string, value string)
	//延迟队列-死信
	PublishDelay(key string, value string, expire string)
	//延迟队列-插件
	PublishXDelay(QueueName string, message string, delayTime int64)
}
