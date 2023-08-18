package dal

import (
	"context"
	"github.com/cold-runner/simpleTikTok/service/user/dal/model"
)

// CreateUser create user info
func CreateUser(ctx context.Context, user []*model.User) error {
	return DB.WithContext(ctx).Create(user).Error
}

// QueryUserByUserName query list of user info
func GetUserByUserName(ctx context.Context,
	userName string) ([]*model.User, error) {
	res := make([]*model.User, 0)
	if err := DB.WithContext(ctx).Where("username =?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
