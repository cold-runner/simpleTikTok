package mq

import (
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/streadway/amqp"
	"time"
)

type UpdateInfoProducer struct {
	channel   *amqp.Channel
	exchange  string
	queueName string
}

type UpdateVideoInfoRequest struct {
	VideoId             int64 `json:"video_id"`
	ActionType          int32 `json:"action_type"`
	UpdateFavoriteCount bool  `json:"update_favorite_count"`
	UpdateCommentCount  bool  `json:"update_comment_count"`
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

func (p *UpdateInfoProducer) Publish(request UpdateVideoInfoRequest) error {
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
