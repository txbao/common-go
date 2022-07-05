package rabbitmq

import (
	"common-go/utils"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/streadway/amqp"
	"log"
	"time"
)

// rabbitmq 的相关配置
const (
	ConsumerName     = ""
	Durable          = true
	DeleteWhenUnused = false
	Exclusive        = false
	NoWait           = false
	AutoAck          = false
	NoLocal          = false
	Mandatory        = false
	Immediate        = false
	DelayExpiration  = "5000" // 设置5秒的队列过期时间, 这里仅仅用在延时队列设置当中

	//exchange格式
	exchangeNameFormat = "%s_exchange"
)

// Rabbitmq 消息队列
type Rabbitmq struct {
	Host     string
	Username string
	Password string
	Vhost    string
}

//获取连接参数
func (rb Rabbitmq) getMqUrl() string {
	password := utils.StrReplace("\\@", "@", rb.Password, -1)

	mqurl := "amqp://" + rb.Username + ":" + password + "@" + rb.Host + ":5672/" + rb.Vhost
	fmt.Println("队列配置:", mqurl)
	return mqurl
}

// Read 向队列读取的方法,消费
func (rb Rabbitmq) Consum(svcCtx interface{}, QueueName string, f func(svcCtx interface{}, jsonStr []byte, msg amqp.Delivery)) {
	defer func() {
		if err := recover(); err != nil {
			time.Sleep(3 * time.Second)
			fmt.Println("Consum_sleep休息3秒")
			rb.Consum(svcCtx, QueueName, f)
		}
	}()
	conn, err := amqp.Dial(rb.getMqUrl())
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明一个主监听队列, 延时队列也将会把过期消息转发到这里
	q, err := ch.QueueDeclare(
		QueueName,        // name
		Durable,          // durable //是否持久化
		DeleteWhenUnused, // delete when unused  是否自动删除
		Exclusive,        // exclusive  //是否具有排他性
		NoWait,           // no-wait //是否阻塞处理
		nil,              // arguments //额外的属性
	)
	failOnError(err, "Failed to declare a queue")

	Exchange := fmt.Sprintf(exchangeNameFormat, QueueName)
	declareDelayQueue(ch, QueueName, Exchange)
	declareExchange(ch, Exchange, "fanout")

	// 将主监听队列和 exchange 绑定
	err = ch.QueueBind(
		q.Name,   // queue name
		"",       // routing key
		Exchange, // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	// start
	closeChan := make(chan *amqp.Error, 1)
	notifyClose := ch.NotifyClose(closeChan) //一旦消费者的channel有错误，产生一个amqp.Error，channel监听并捕捉到这个错误
	closeFlag := false
	// end

	msgs, err := ch.Consume(
		q.Name,       // queue
		ConsumerName, // consumer //用来区分多个消费者
		AutoAck,      // auto-ack //是否自动应答
		Exclusive,    // exclusive //是否独有
		NoLocal,      // no-local  //设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		NoWait,       // no-wait  //列是否阻塞
		nil,          // args
	)
	failOnError(err, "Failed to register a consumer")

	//start 消费者模式 断线重连机制
	//var obj Obj
	for {
		select {
		case e := <-notifyClose:
			fmt.Println("chan通道错误,e:%s", e.Error())
			close(closeChan)
			time.Sleep(5 * time.Second)
			rb.Consum(svcCtx, QueueName, f)
			closeFlag = true
		case msg := <-msgs:
			log.Printf("MQ消费接收，Received a message: %s,时间：%s", msg.Body, utils.DateNowFormatStr())
			go f(svcCtx, msg.Body, msg)
			log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		}
		if closeFlag {
			break
		}
	}
	//end
	/*
		//老方法
		forever := make(chan bool)

		go func() {
			for d := range msgs {
				log.Printf("MQ消费接收，Received a message: %s,时间：%s", d.Body, utils.DateNowFormatStr())
				go f(d.Body)
			}
		}()

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<-forever
	*/
}


//发布延迟队列-插件
func (rb Rabbitmq) PublishXDelay(QueueName string, message string, delayTime int64) {
	defer func() {
		if err := recover(); err != nil {
			time.Sleep(3 * time.Second)
			fmt.Println("PublishXDelay_sleep休息3秒")
			rb.PublishXDelay(QueueName, message, delayTime)
		}
	}()
	conn, err := amqp.Dial(rb.getMqUrl())
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	_ = ch.ExchangeDeclare(QueueName+".delayed", "x-delayed-message",
		true, false, false, false, amqp.Table{"x-delayed-type": "direct"})
	_ = ch.QueueBind(QueueName, QueueName, QueueName+".delayed", false, nil)
	//调用channel 发送消息到队列中
	err = ch.Publish(
		QueueName+".delayed",
		QueueName,
		//如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(message),
			DeliveryMode: 2,
			Headers:      amqp.Table{"x-delay": delayTime * 1000},
		})
	if err != nil {
		fmt.Println("发布延迟队列有误,err:%s", err)
	}
}

// Delay 发送延时消息-死信
func (rb Rabbitmq) PublishDelay(QueueName string, key string, value string, expire int64) {
	defer func() {
		if err := recover(); err != nil {
			time.Sleep(3 * time.Second)
			fmt.Println("PublishDelay_sleep休息3秒")
			rb.PublishDelay(QueueName, key, value, expire)
		}
	}()
	conn, err := amqp.Dial(rb.getMqUrl())
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	body := value

	delayName := fmt.Sprintf("%s_delay", QueueName)

	expireStr := decimal.NewFromInt(expire).Mul(decimal.NewFromFloat(1000)).String()
	err = ch.Publish(
		"",        // exchange
		delayName, // routing key
		Mandatory, // mandatory
		Immediate, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
			Expiration:  expireStr, // 设置五秒的过期时间
		})
	log.Printf("延迟队列 [x] DelayMQSent %s,时间:%v秒,MQ名称：%s", body, expire, QueueName)
	failOnError(err, "Failed to publish a message")
	if err != nil {
		time.Sleep(3 * time.Second)
		rb.PublishDelay(QueueName, key, value, expire)
	}
}

// 声明一个延时队列,这个队列不做消费,而是让消息变成死信后再进行转发
func declareDelayQueue(ch *amqp.Channel, channelName string, exchangeName string) {
	delayName := channelName + "_delay"
	_, errDelay := ch.QueueDeclare(
		delayName, // name
		false,     // durable //是否持久化
		false,     // delete when unused 是否自动删除
		false,     // exclusive //是否具有排他性
		false,     // no-wait  //是否阻塞处理
		amqp.Table{
			"x-dead-letter-exchange": exchangeName,
		}, // arguments //额外的属性
	)
	failOnError(errDelay, "Failed to declare a delay_queue")
}

// 声明一个 exchange, 这里只是为了接收延时队列而设置的一个 exchange
func declareExchange(ch *amqp.Channel, exchangeName string, exType string) {
	err := ch.ExchangeDeclare(
		exchangeName, // name
		exType,       // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare an exchange")
}

// Send 向队列发送的方法
func (rb Rabbitmq) Publish(QueueName string, key string, value string) {
	defer func() {
		if err := recover(); err != nil {
			time.Sleep(3 * time.Second)
			fmt.Println("Publish_sleep休息3秒")
			rb.Publish(QueueName, key, value)
		}
	}()
	conn, err := amqp.Dial(rb.getMqUrl())
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		QueueName,        // name
		Durable,          // durable
		DeleteWhenUnused, // delete when unused
		Exclusive,        // exclusive
		NoWait,           // no-wait
		nil,              // arguments
	)
	failOnError(err, "Failed to declare a queue")

	exchange := fmt.Sprintf(exchangeNameFormat, QueueName)
	body := value
	err = ch.Publish(
		exchange,  // exchange
		q.Name,    // routing key
		Mandatory, // mandatory
		Immediate, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf("普通队列 [x] MQSent %s", body)
	failOnError(err, "Failed to publish a message")
	if err != nil {
		time.Sleep(3 * time.Second)
		rb.Publish(QueueName, key, value)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		//log.Fatalf("%s: %s", msg, err)
		fmt.Println(fmt.Sprintf("%s: %s", msg, err))
	}
}
