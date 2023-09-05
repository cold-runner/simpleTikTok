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

func FavoriteList(ctx context.Context,
	req *SocialService.FavoriteListRequest) (*SocialService.FavoriteListResponse, error) {
	log.Debugw("Calling rpc favorite list request:", "req", req)

	resp, err := socialClient.FavoriteList(ctx, req)
	if err != nil {
		log.Errorw("calling api server rpc favorite list failed", "err", err)
		return nil, err
	}
	log.Debugw("api server rpc favorite list success", "resp", resp)
	return resp, nil
}

func CommentAction(ctx context.Context,
	req *SocialService.CommentActionRequest) (*SocialService.
	CommentActionResposne, error) {
	log.Debugw("Calling rpc comment action request:", "req", req)
	resp, err := socialClient.CommentAction(ctx, req)
	if err != nil {
		log.Errorw("calling api server rpc comment action failed", "err", err)
		return nil, err
	}
	log.Debugw("api server rpc comment action success", "resp", resp)
	return resp, nil
}

func CommentList(ctx context.Context, req *SocialService.CommentListRequest) (
	*SocialService.CommentListResponse, error) {
	log.Debugw("Calling rpc comment list request:", "req", req)
	resp, err := socialClient.CommentList(ctx, req)
	if err != nil {
		log.Errorw("calling api server rpc comment list failed", "err", err)
		return nil, err
	}
	log.Debugw("api server rpc comment list success", "resp", resp)
	return resp, nil
}
