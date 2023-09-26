package mq

import (
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

var QueueName = "userInfoUpdate"
var UserInfoUpdateProducer *Producer

var SocialQueueName = "SocialInfoUpdate"
var SocialInfoUpdateProducer *UpdateInfoProducer

func InitMQ() {
	conn, err := amqp.Dial(viper.GetString("mq.url"))
	if err != nil {
		log.Panicw("connect to mq failed", "err", err)
	}
	// 创建生产者
	UserInfoUpdateProducer, err = NewProducer(conn, QueueName)
	if err != nil {
		log.Panicw("init mq producer failed", "err", err)
	}

	// 开启消费监听
	relationConsumer, err := NewConsumer(conn, QueueName)
	if err != nil {
		log.Fatalw("init mq consumer failed", "err", err)
	}
	go relationConsumer.Consume()

}

func InitSocialMQ() {
	conn, err := amqp.Dial(viper.GetString("mq.url"))
	if err != nil {
		log.Panicw("connect to mq failed", "err", err)
	}
	// 创建生产者
	SocialInfoUpdateProducer, err = NewUpdateInfoProducer(conn,
		SocialQueueName)
	if err != nil {
		log.Panicw("init mq producer failed", "err", err)
	}

	// 开启消费监听
	SocialInfoConsumer, err := NewSocialServiceConsumer(conn, SocialQueueName)
	if err != nil {
		log.Fatalw("init mq consumer failed", "err", err)
	}
	go SocialInfoConsumer.Consume()

}
