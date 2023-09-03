package mq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"time"
)

type VideoPublishRequest struct {
	AuthorID int64  `json:"author_id"`
	Data     []byte `json:"data"`
	Title    string `json:"title"`
}

// Producer 生产者结构体
type Producer struct {
	channel   *amqp.Channel
	exchange  string
	queueName string
}

func NewProducer(conn *amqp.Connection, queueName string) (*Producer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &Producer{
		channel:   ch,
		queueName: queueName,
	}, nil
}

func (p *Producer) Publish(request VideoPublishRequest) error {
	message, err := json.Marshal(request)
	if err != nil {
		return err
	}

	// 队列声明
	_, err = p.channel.QueueDeclare(p.queueName, false, false, false, false, nil)
	if err != nil {
		return err
	}

	// 发布消息
	return p.channel.Publish(p.exchange, p.queueName, false, false, amqp.Publishing{
		Timestamp:   time.Now(),
		ContentType: "application/json",
		Body:        message,
	})
}
