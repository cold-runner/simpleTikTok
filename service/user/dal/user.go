package dal

import (
	"context"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/user/dal/model"
)

// CreateUser create user info
func CreateUser(ctx context.Context, user *model.User) error {
	err := DB.WithContext(ctx).Create(user).Error
	if err != nil {
		log.Errorw("create user failed", "err", err)
		return err
	}
	return nil
}

// GetUserByUserName one to one 获取 User 对象
func GetUserByUserName(ctx context.Context,
	userName string) (*model.User, error) {
	res := &model.User{}
	if err := DB.WithContext(ctx).Where("username = ?",
		userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
