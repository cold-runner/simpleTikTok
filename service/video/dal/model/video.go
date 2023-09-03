package model

import (
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"gorm.io/gorm"
	"time"
)

const TableNameVideo = "video"

type Video struct {
	VideoID       int64     `gorm:"column:video_id;type:int;primaryKey" json:"video_id"`
	AuthorID      int64     `gorm:"column:author_id;type:int;" json:"author_id"`
	PlayUrl       string    `gorm:"column:play_url;type:varchar(255);not null" json:"play_url"`
	CoverUrl      string    `gorm:"column:cover_url;type:varchar(255);not null" json:"cover_url"`
	FavoriteCount int64     `gorm:"column:favorite_count;type:int;default:0;not null" json:"favorite_count"`
	CommentCount  int64     `gorm:"column:comment_count;type:int;default:0;not null" json:"comment_count"`
	Title         string    `gorm:"column:title;type:varchar(255);not null" json:"title"`
	CreatedAt     time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;not null"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:timestamp;default
:CURRENT_TIMESTAMP on update current_timestamp;not null"`
}

func (*Video) TableName() string {
	return TableNameVideo
}

// BeforeCreate 在创建数据库记录之前，设置记录的创建时间
func (v *Video) BeforeCreate(tx *gorm.DB) (err error) {
	v.CreatedAt = time.Now().Local()
	v.UpdatedAt = time.Now().Local()
	if err != nil {
		return err
	}
	return nil
}

func (v *Video) BeforeUpdate(tx *gorm.DB) error {
	v.UpdatedAt = time.Now().Local()
	return nil
}

func (v *Video) validate(tx *gorm.DB) error {
	if v.FavoriteCount < 0 || v.CommentCount < 0 {
		log.Errorw("invalid count", "video", v)
		return errno.ErrInvalidUpdate
	}
	return nil
}
