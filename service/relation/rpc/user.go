package rpc

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/BaseResponse"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/pkg/log"
)

func UpdateUserFollowCount(ctx context.Context, req *UserService.ChangeUserFollowCountRequest) (*BaseResponse.BaseResp,
	error) {

	resp, err := userClient.ChangeUserFollowCount(ctx, req)
	if err != nil {
		log.Errorw("api server rpc change user follow count failed", "err", err)
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErr(resp.StatusCode,
			"", resp.StatusMsg)
	}
	return resp, nil
}

func MGetUserInfo(ctx context.Context,
	req *UserService.MGetUserInfoRequest) (*UserService.
	MGetUserInfoResponse, error) {
	resp, err := userClient.MGetUserInfo(ctx, req)
	if err != nil {
		log.Errorw("api server rpc m get user info failed", "err", err)
	}
	if resp.GetBaseResp().GetStatusCode() != 0 {
		return nil, errno.NewErr(resp.GetBaseResp().GetStatusCode(),
			"", resp.GetBaseResp().GetStatusMsg())
	}
	return resp, nil
}
