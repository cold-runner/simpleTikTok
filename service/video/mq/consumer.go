package mq

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/video/dal"
	"github.com/cold-runner/simpleTikTok/service/video/dal/model"
	"github.com/cold-runner/simpleTikTok/service/video/minio"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"sync"
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

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			var request VideoPublishRequest
			if err := json.Unmarshal(msg.Body, &request); err != nil {
				log.Errorw("Cannot unmarshal message", "err", err)
				continue
			}
			videoUUID, err := uuid.NewRandom()
			if err != nil {
				log.Errorw("generate video uuid failed", "err", err)
			}
			videoName := videoUUID.String() + "." + "mp4"
			imageName := videoUUID.String() + "." + "jpg"
			videoURL := minio.BuildVideoURL(videoName)
			imageURL := minio.BuildImageURL(imageName)
			var wg sync.WaitGroup
			// 上传视频到Minio
			wg.Add(1)
			go func() {
				defer wg.Done()
				err := minio.UploadVideo(context.Background(), videoName, request.Data)
				if err != nil {
					log.Errorw("Failed to upload to Minio", "err", err)
				}
			}()

			// 上传视频数据到MySQL
			wg.Add(1)
			go func() {
				defer wg.Done()
				videoInfo := &model.Video{
					AuthorID:      request.AuthorID,
					PlayUrl:       videoURL,
					CoverUrl:      imageURL,
					FavoriteCount: 0,
					CommentCount:  0,
					Title:         request.Title,
				}
				err := dal.CreateVideo(context.Background(), videoInfo)
				if err != nil {
					log.Errorw("Failed to upload metadata to MySQL", "err", err)
				}
			}()
			wg.Wait()

			// 进行封面截取并上传
			coverBuf, err := minio.ExtractCoverByVideoURL(videoURL)
			if err != nil {
				log.Errorw("Failed to extract cover", "err", err)
			}
			err = minio.UploadImage(context.Background(), imageName, coverBuf)
			if err != nil {
				log.Errorw("Failed to upload cover to Minio", "err", err)
			}
			log.Debugw("Success to publish video", "videoName", videoName, "imageName", imageName)
		}
	}()

	<-forever
}
