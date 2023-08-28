package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/pkg/encryption"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/service/user/dal"
)

type ValidateLoginService struct {
	ctx context.Context
}

func NewValidateLoginService(ctx context.Context) *ValidateLoginService {
	return &ValidateLoginService{ctx: ctx}
}

func (s *ValidateLoginService) ValidateLogin(req *UserService.
	UserLoginRequest) (int64, error) {
	u, err := dal.GetUserByUserName(s.ctx, req.Username)
	// 判断用户是否存在
	if u == nil {
		return -1, errno.ErrUserNotExist
	}
	// 判断获取用户信息是否出错
	if err != nil {
		return -1, err
	}
	// 判断密码是否正确
	err = encryption.Compare(u.Password, req.Password)
	if err != nil {
		return -1, errno.ErrUserPassword
	}

	return u.ID, nil
}
