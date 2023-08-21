package rpc

import "github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"

type UserRpc interface {
	GetUserInfo(int64) (*SocialService.User, error)
}
