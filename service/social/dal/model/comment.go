package model

import (
	"gorm.io/gorm"
	"time"
)

const TableNameComment = "video_comment"

type VideoComment struct {
	commentId       int64     `gorm:"primaryKey;column:comment_id"`
	commentVideoId  int64     `gorm:"column:comment_video_id"`
	commentAuthorId int64     `gorm:"column:comment_author_id"`
	commentContent  string    `gorm:"column:comment_content"`
	CreatedAt       time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;not null"`
}

// BeforeCreate 在创建数据库记录之前，设置记录的创建时间
func (c *VideoComment) BeforeCreate(tx *gorm.DB) (err error) {
	c.CreatedAt = time.Now().Local()
	if err != nil {
		return err
	}
	return nil
}
