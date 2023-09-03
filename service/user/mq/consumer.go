package mq

import (
	"context"
	"encoding/json"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/user/dal"
	"github.com/streadway/amqp"
)

// Consumer 消费者结构体
type Consumer struct {
	channel   *amqp.Channel
	queueName string
}

type SocialServiceConsumer struct {
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

func NewSocialServiceConsumer(conn *amqp.Connection, queueName string) (*SocialServiceConsumer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &SocialServiceConsumer{
		channel:   ch,
		queueName: queueName,
	}, nil
}

func (c *SocialServiceConsumer) Consume() {
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
		var request UpdateUserInfoRequest
		if err := json.Unmarshal(msg.Body, &request); err != nil {
			log.Errorw("failed to unmarshal message: %v", "err", err)
			continue
		}
		ctx := context.Background()
		log.Debugw("mq start update user info", "request", request)
		// 处理用户信息更新
		if request.UpdateTotalFavorited {
			// 更新用户获赞数量
			if request.ActionType == 1 {
				err = dal.IncreaseUserTotalFavorited(ctx,
					request.ToUserId)
			} else if request.ActionType == 2 {
				err = dal.DecreaseUserTotalFavorited(ctx,
					request.ToUserId)
			}
			log.Debugw("mq update user total favorited success")

		} else if request.UpdateWorkCount {
			// 更新用户作品数量
			if request.ActionType == 1 {
				err = dal.IncreaseUserWorkCount(ctx,
					request.ToUserId)
			} else if request.ActionType == 2 {
				err = dal.DecreaseUserWorkCount(ctx, request.ToUserId)
			}
			log.Debugw("mq update user work count success")

		} else if request.UpdateFavoriteCount {
			// 更新用户喜欢数
			if request.ActionType == 1 {
				err = dal.IncreaseUserFavoriteCount(ctx, request.ToUserId)
			} else if request.ActionType == 2 {
				err = dal.DecreaseUserFavoriteCount(ctx, request.ToUserId)
			}
			log.Debugw("mq update user favorite count success")
		}
		if err != nil {
			log.Errorw("failed to update user info", "err", err)
		}

	}
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
		var request ChangeUserFollowCountRequest
		if err := json.Unmarshal(msg.Body, &request); err != nil {
			log.Errorw("failed to unmarshal message: %v", "err", err)
			continue
		}

		// 处理关注/取关操作的粉丝和关注数量更新
		if request.ActionType == 1 {
			err = followActionCountUpdate(request.UserId, request.ToUserId)
			if err != nil {
				log.Errorw("failed to update follow action info", "err", err)
			}
		} else if request.ActionType == 2 {
			err = unfollowActionCountUpdate(request.UserId, request.ToUserId)
			if err != nil {
				log.Errorw("failed to update unfollow action info", "err", err)
			}
		}
	}
}

func followActionCountUpdate(uid int64, toUid int64) error {
	err := dal.IncreaseUserFollowCount(context.Background(), uid)
	if err != nil {
		return err
	}
	err = dal.IncreaseUserFollowerCount(context.Background(), toUid)
	if err != nil {
		return err
	}
	return nil
}

func unfollowActionCountUpdate(uid int64, toUid int64) error {
	err := dal.DecreaseUserFollowCount(context.Background(), uid)
	if err != nil {
		return err
	}
	err = dal.DecreaseUserFollowerCount(context.Background(), toUid)
	if err != nil {
		return err
	}
	return nil
}
