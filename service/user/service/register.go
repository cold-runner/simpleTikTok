package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/service/user/dal"
	"github.com/cold-runner/simpleTikTok/service/user/dal/model"
)

type CreateUserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// CreateUser 执行创建新用户流程.
func (s *CreateUserService) CreateUser(req *UserService.
	UserRegisterRequest) error {
	u, err := dal.GetUserByUserName(s.ctx, req.Username)
	if err != nil {
		return err
	}
	// 检查用户名是否已经存在
	if req.Username == u.Username {
		return errno.ErrUserAlredyExist
	}
	// model.user 中的BeforeCreate() Hook 方法会执行密码加密操作，而不会使用明文加密，
	// 这里便不做处理，直接创建数据库记录
	return dal.CreateUser(s.ctx, &model.User{
		Username: req.Username,
		Password: req.Password,
	})
}
