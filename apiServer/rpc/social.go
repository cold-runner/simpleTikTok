package rpc

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/SocialService/socialservice"
	"github.com/cold-runner/simpleTikTok/pkg/log"
)

var socialClient socialservice.Client

func FavoriteAction(ctx context.Context,
	req *SocialService.FavoriteActionRequest) (*SocialService.
	FavoriteActionResponse, error) {
	log.Debugw("Calling rpc favorite action request:", "req", req)

	resp, err := socialClient.FavoriteAction(ctx, req)
	if err != nil {
		log.Errorw("calling api server rpc favorite action failed", "err", err)
		return nil, err
	}
	log.Debugw("api server rpc favorite action success", "resp", resp)
	return resp, nil
}
