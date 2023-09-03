package model

import (
	"github.com/cold-runner/simpleTikTok/pkg/encryption"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"gorm.io/gorm"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID              int64  `gorm:"column:id;type:int;primaryKey" json:"id,string"`
	Username        string `gorm:"column:username;type:varchar(32)" json:"username"`
	Password        string `gorm:"column:password;type:varchar(255)" json:"password"`
	FollowCount     int64  `gorm:"column:follow_count;type:bigint" json:"follow_count"`
	FollowerCount   int64  `gorm:"column:follower_count;type:bigint" json:"follower_count"`
	Avatar          string `gorm:"column:avatar;type:varchar(255)" json:"avatar"`
	BackgroundImage string `gorm:"column:background_image;type:varchar(255)" json:"background_image"`
	Signature       string `gorm:"column:signature;type:text" json:"signature"`
	TotalFavorited  int64  `gorm:"column:total_favorited;type:bigint" json:"total_favorited"`
	WorkCount       int64  `gorm:"column:work_count;type:bigint" json:"work_count"`
	FavoriteCount   int64  `gorm:"column:favorite_count;type:bigint" json:"favorite_count"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}

// BeforeCreate 在创建数据库记录之前加密明文密码.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password, err = encryption.Encrypt(u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	return u.validateCounts()
}

// validateCounts 验证所有的Count字段都必须 >= 0
func (u *User) validateCounts() error {
	if u.FollowCount < 0 || u.FollowerCount < 0 || u.TotalFavorited < 0 || u.WorkCount < 0 || u.FavoriteCount < 0 {
		log.Errorw("invalid count", "user", u)
		return errno.ErrInvalidUpdate
	}
	return nil
}
