package mq

import (
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

var RelationQueueName = "relation"
var RelationProducer *Producer

func InitMQ() {
	conn, err := amqp.Dial(viper.GetString("mq.url"))
	if err != nil {
		log.Panicw("connect to mq failed", "err", err)
	}
	// 创建生产者
	RelationProducer, err = NewProducer(conn, RelationQueueName)
	if err != nil {
		log.Panicw("init mq producer failed", "err", err)
	}

	// 开启消费监听
	relationConsumer, err := NewConsumer(conn, RelationQueueName)
	if err != nil {
		log.Fatalw("init mq consumer failed", "err", err)
	}
	go relationConsumer.Consume()

}
