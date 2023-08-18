package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/service/user/dal"
)

type CreateUserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// CreateUser create user
func (s *CreateUserService) CreateUser(req *UserService.
	UserRegisterRequest) error {
	u, err := dal.GetUserByUserName(s.ctx, req.Username)
	if err != nil {
		return
	}

	if req.Username == u.Username {
		return errno.ErrUserAlredyExist
	}

}
