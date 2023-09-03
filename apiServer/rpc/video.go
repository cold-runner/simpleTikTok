package rpc

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/VideoService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/VideoService/videoservice"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/pkg/log"
)

var videoClient videoservice.Client

func PublishVideo(ctx context.Context,
	req *VideoService.VideoPublishActionRequest) (*VideoService.
	VideoPublishActionResponse, error) {
	log.Debugw("Calling rpc publish video request:", "Author",
		req.UserId, "Video Title:",
		req.Title)
	resp, err := videoClient.VideoPublishAction(ctx, req)
	if err != nil {
		log.Errorw("calling api server rpc publish video failed", "err", err)
		return nil, err
	}
	if resp.BaseResp.GetStatusCode() != 0 {
		log.Errorw("response error", "resp", resp)
		return nil, errno.NewErr(resp.BaseResp.GetStatusCode(), "",
			resp.BaseResp.GetStatusMsg())
	}
	log.Debugw("api server rpc publish video success", "resp", resp)
	return resp, nil
}

func VideoFeed(ctx context.Context, req *VideoService.VideoFeedRequest) (
	*VideoService.VideoFeedResponse, error) {
	log.Debugw("Calling rpc video feed request:", "req", req)
	resp, err := videoClient.VideoFeed(ctx, req)
	if err != nil {
		log.Errorw("calling api server rpc video feed failed", "err", err)
		return nil, err
	}
	if resp.BaseResp.GetStatusCode() != 0 {
		log.Errorw("response error", "resp", resp)
		return nil, errno.NewErr(resp.BaseResp.GetStatusCode(), "",
			resp.BaseResp.GetStatusMsg())
	}
	log.Debugw("api server rpc video feed success", "resp", resp)
	return resp, nil
}

func VideoPublishList(ctx context.Context, req *VideoService.VideoPublishListRequest) (
	*VideoService.VideoPublishListResponse, error) {
	log.Debugw("Calling rpc video publish list request:", "req", req)
	resp, err := videoClient.VideoPublishList(ctx, req)
	if err != nil {
		log.Errorw("calling api server rpc video publish list failed", "err", err)
		return nil, err
	}
	if resp.BaseResp.GetStatusCode() != 0 {
		log.Errorw("response error", "resp", resp)
		return nil, errno.NewErr(resp.BaseResp.GetStatusCode(), "",
			resp.BaseResp.GetStatusMsg())
	}
	log.Debugw("api server rpc video publish list success", "resp", resp)
	return resp, nil
}
