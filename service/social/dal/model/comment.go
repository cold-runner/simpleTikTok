package model

import (
	"gorm.io/gorm"
	"time"
)

const TableNameComment = "video_comment"

type VideoComment struct {
	CommentId int64  `gorm:"primaryKey;column:comment_id"`
	VideoId   int64  `gorm:"column:video_id"`
	AuthorId  int64  `gorm:"column:author_id"`
	Content   string `gorm:"column:content"`
	CreatedAt string `gorm:"column:created_at;type:varchar(5)"`
}

// BeforeCreate 在创建数据库记录之前，设置记录的创建时间
func (c *VideoComment) BeforeCreate(tx *gorm.DB) (err error) {
	c.CreatedAt = time.Now().Format("01-02")
	if err != nil {
		return err
	}
	return nil
}
func (*VideoComment) TableName() string {
	return TableNameComment
}
