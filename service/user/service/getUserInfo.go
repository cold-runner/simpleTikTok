package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/service/user/dal"
	"github.com/cold-runner/simpleTikTok/service/user/dal/model"
)

type GetUserInfoService struct {
	ctx context.Context
}

// NewGetUserInfoService create new GetUserInfoService
func NewGetUserInfoService(ctx context.Context) *GetUserInfoService {
	return &GetUserInfoService{ctx: ctx}
}

// GetUserInfoByName get user info by username
func (s *GetUserInfoService) GetUserInfoByName(username string) (*model.User,
	error) {
	u, err := dal.GetUserByUserName(s.ctx, username)
	if err != nil {
		return u, errno.ErrUserNotExist
	}
	return u, nil
}
