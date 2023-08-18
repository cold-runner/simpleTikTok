package main

import (
	"context"
	UserService "github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *UserService.UserRegisterRequest) (resp *UserService.UserRegisterResponse, err error) {
	// TODO: Your code here...

	return nil, nil
}
