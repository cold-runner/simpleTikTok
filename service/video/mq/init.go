package mq

import (
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

var VideoQueueName = "video"
var VideoInfoQueueName = "videoInfo"
var VideoProducer *Producer
var VideoInfoProducer *UpdateInfoProducer

func InitMQ() {
	conn, err := amqp.Dial(viper.GetString("mq.url"))
	if err != nil {
		log.Panicw("connect to mq failed", "err", err)
	}
	// 创建生产者
	VideoProducer, err = NewProducer(conn, VideoQueueName)
	VideoInfoProducer, err = NewUpdateInfoProducer(conn, VideoInfoQueueName)
	if err != nil {
		log.Panicw("init mq producer failed", "err", err)
	}

	// 开启消费监听
	videoConsumer, err := NewConsumer(conn, VideoQueueName)
	if err != nil {
		log.Fatalw("init mq consumer failed", "err", err)
	}
	videoInfoConsumer, err := NewConsumer(conn, VideoInfoQueueName)

	if err != nil {
		log.Fatalw("init mq consumer failed", "err", err)
	}
	go videoConsumer.Consume()
	go videoInfoConsumer.ConsumeUpdateRequest()
}
