package mq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"time"
)

// ChangeUserFollowCountRequest 关注/取关请求修改信息请求结构体
type ChangeUserFollowCountRequest struct {
	UserId     int64 `json:"user_id"`
	ToUserId   int64 `json:"to_user_id"`
	ActionType int32 `json:"action_type"`
}

type UpdateUserInfoRequest struct {
	ToUserId             int64 `json:"to_user_id"`
	ActionType           int32 `json:"action_type"`
	UpdateTotalFavorited bool  `json:"update_total_favorited"`
	UpdateWorkCount      bool  `json:"update_work_count"`
	UpdateFavoriteCount  bool  `json:"update_favorite_count"`
}

// Producer 生产者结构体
type Producer struct {
	channel   *amqp.Channel
	exchange  string
	queueName string
}

type UpdateInfoProducer struct {
	channel   *amqp.Channel
	exchange  string
	queueName string
}

func NewUpdateInfoProducer(conn *amqp.Connection,
	queueName string) (*UpdateInfoProducer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &UpdateInfoProducer{
		channel:   ch,
		queueName: queueName,
	}, nil
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

func (p *Producer) Publish(request ChangeUserFollowCountRequest) error {
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

func (p *UpdateInfoProducer) Publish(request UpdateUserInfoRequest) error {
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
