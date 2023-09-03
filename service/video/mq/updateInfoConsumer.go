package mq

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/video/dal"
)

func (c *Consumer) ConsumeUpdateRequest() {
	// 队列声明
	_, err := c.channel.QueueDeclare(c.queueName, false, false, false, false, nil)
	if err != nil {
		log.Fatalw("cannot declare queue", "err", err)
	}

	msgs, err := c.channel.Consume(c.queueName, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalw("cannot Consume queue", "err", err)
	}

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			var request UpdateVideoInfoRequest
			if err := json.Unmarshal(msg.Body, &request); err != nil {
				log.Errorw("Cannot unmarshal message", "err", err)
				continue
			}
			ctx := context.Background()

			if request.UpdateFavoriteCount {
				switch request.ActionType {
				case 1:
					err = dal.IncrementFavoriteCount(ctx, request.VideoId)
					if err != nil {
						log.Errorw("increment FavoriteCount failed", "err", err)
					}
				case 2:
					err = dal.DecrementFavoriteCount(ctx, request.VideoId)
					if err != nil {
						log.Errorw("decrement FavoriteCount failed", "err", err)
					}
				}

			} else if request.UpdateCommentCount {
				switch request.ActionType {
				case 1:
					err = dal.IncrementCommentCount(ctx, request.VideoId)
					if err != nil {
						log.Errorw("increment CommentCount failed", "err", err)
					}
				case 2:
					err = dal.DecrementCommentCount(ctx, request.VideoId)
					if err != nil {
						log.Errorw("decrement CommentCount failed", "err", err)
					}
				}
			}
		}
	}()

	<-forever
}
