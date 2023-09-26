package rpc

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/RelationService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/RelationService/relationservice"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/pkg/log"
)

var relationClient relationservice.Client

func RelationAction(ctx context.Context, request *RelationService.
	RelationActionRequest) (*RelationService.RelationActionResponse, error) {
	log.Infow("Calling rpc relation action request:", "req", request)
	resp, err := relationClient.RelationAction(ctx, request)
	// 处理调用rpc 服务时可能发生的错误
	if err != nil {
		log.Errorw("api server rpc relation action failed", "err", err)
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErr(resp.BaseResp.StatusCode,
			"", resp.BaseResp.StatusMsg)
	}
	return resp, nil

}

func GetFollowList(ctx context.Context, request *RelationService.
	RelationFollowListRequest) (*RelationService.RelationFollowListResponse,
	error) {
	log.Infow("Calling rpc get follow list request:", "req", request)
	resp, err := relationClient.GetFollowList(ctx, request)
	log.Debugw("api server rpc get follow list success", "resp", resp)
	// 处理调用rpc 服务时可能发生的错误
	if err != nil {
		log.Errorw("api server rpc get follow list failed", "err", err)
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErr(resp.BaseResp.StatusCode,
			"", resp.BaseResp.StatusMsg)
	}
	return resp, nil
}

func GetFollowerList(ctx context.Context,
	requset *RelationService.RelationFollowerListRequest) (*RelationService.
	RelationFollowerListResponse, error) {
	log.Infow("Calling rpc get follower list request:", "req", requset)
	resp, err := relationClient.GetFollowerList(ctx, requset)
	log.Debugw("api server rpc get follower list success", "resp", resp)

	if err != nil {
		log.Errorw("api server rpc get follower list failed", "err", err)
		return nil, err
	}
	if resp.BaseResp.GetStatusCode() != 0 {
		return nil, errno.NewErr(resp.BaseResp.GetStatusCode(),
			"", resp.BaseResp.GetStatusMsg())
	}
	return resp, nil

}
