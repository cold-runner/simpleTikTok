package mq

import (
	"context"
	"encoding/json"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/relation/dal"
	"github.com/streadway/amqp"
)

// Consumer 消费者结构体
type Consumer struct {
	channel   *amqp.Channel
	queueName string
}

func NewConsumer(conn *amqp.Connection, queueName string) (*Consumer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &Consumer{
		channel:   ch,
		queueName: queueName,
	}, nil
}

func (c *Consumer) Consume() {
	// 队列声明
	_, err := c.channel.QueueDeclare(c.queueName, false, false, false, false, nil)
	if err != nil {
		log.Fatalw("cannot declare queue", "err", err)
	}

	msgs, err := c.channel.Consume(c.queueName, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalw("cannot Consume queue", "err", err)
	}

	for msg := range msgs {
		var request RelationActionRequest
		if err := json.Unmarshal(msg.Body, &request); err != nil {
			log.Errorw("failed to unmarshal message: %v", "err", err)
			continue
		}

		// 处理关注/取关操作
		// 此处可以执行数据库操作，伪代码示例如下：
		if request.ActionType == 1 {
			// 执行关注操作，更新MySQL和Redis
			err = followAction(request.UserId, request.ToUserId)
			if err != nil {
				log.Errorw("failed to execute follow action", "err", err)
			}
		} else if request.ActionType == 2 {
			//     // 执行取关操作，更新MySQL和Redis
			err = unfollowAction(request.UserId, request.ToUserId)
			if err != nil {
				log.Errorw("failed to execute unfollow action", "err", err)
			}
		}
	}
}

func followAction(uid, toUid int64) error {
	ctx := context.Background()
	err := dal.CreateFollowRelation(ctx, uid, toUid)
	if err != nil {
		log.Errorw("RelationActionService.followAction", "err", err)
		return err
	}
	log.Infow("Success to follow", "uid", uid, "toUid", toUid)
	return nil
}

func unfollowAction(uid, toUid int64) error {
	// 判断是否已经关注
	ctx := context.Background()
	err := dal.DeleteFollowRelation(ctx, uid, toUid)
	if err != nil {
		log.Errorw("RelationActionService.unfollowAction", "err", err)
	}
	log.Infow("Success to unfollow", "uid", uid, "toUid", toUid)
	return nil
}
