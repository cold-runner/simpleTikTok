package main

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/VideoService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/service/video/response"
	"github.com/cold-runner/simpleTikTok/service/video/service"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// VideoFeed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoFeed(ctx context.Context, req *VideoService.VideoFeedRequest) (resp *VideoService.VideoFeedResponse, err error) {
	// TODO: Your code here...
	videoList, latestTime, err := service.NewVideoFeedService(ctx).VideoFeed(req)
	if err != nil {
		resp = response.BuildVideoFeedResp(nil, 0, err)
		return nil, err
	}
	resp = response.BuildVideoFeedResp(videoList, latestTime, errno.OK)
	return resp, nil
}

// VideoPublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoPublishAction(ctx context.Context, req *VideoService.VideoPublishActionRequest) (resp *VideoService.VideoPublishActionResponse, err error) {
	// TODO: Your code here...
	if req.UserId <= 0 || len(req.Data) < 0 || len(req.Title) < 0 {
		return nil, errno.ErrInvalidParameter
	}

	err = service.NewPublishActionService(ctx).PublishAction(req)
	if err != nil {
		resp = response.BuildPublishActionResp(err)
		return resp, err
	}
	resp = response.BuildPublishActionResp(errno.OK)
	return resp, nil

}

// VideoPublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoPublishList(ctx context.Context, req *VideoService.VideoPublishListRequest) (resp *VideoService.VideoPublishListResponse, err error) {
	// TODO: Your code here...
	if req.UserId <= 0 || req.ToUserId <= 0 {
		return nil, errno.ErrInvalidParameter
	}

	videoList, err := service.NewVideoPublishListService(ctx).GetPublishVideoList(req)
	if err != nil {
		resp = response.BuildVideoListResp(nil, err)
		return resp, err
	}
	resp = response.BuildVideoListResp(videoList, errno.OK)
	return resp, err
}

// VideoPublishListByIds implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoPublishListByIds(ctx context.Context, req *VideoService.VideoPublishListByIdsRequest) (resp *VideoService.VideoPublishListResponse, err error) {
	// TODO: Your code here...
	if req.GetUserId() <= 0 || len(req.GetToVideoIds()) < 0 {
		return nil, errno.ErrInvalidParameter
	}

	videoList, err := service.NewVideoListService(ctx).GetVideoListByIds(req)
	if err != nil {
		resp = response.BuildVideoListResp(nil, err)
		return resp, err

	}
	resp = response.BuildVideoListResp(videoList, errno.OK)
	return resp, err
}

// UpdateVideoFavoriteCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) UpdateVideoFavoriteCount(ctx context.Context, req *VideoService.UpdateVideoFavoriteCountRequest) (resp *VideoService.UpdateVideoInfoResponse, err error) {
	// TODO: Your code here...
	err = service.NewUpdateVideoInfoService(ctx).UpdateVideoFavoriteCount(req)
	if err != nil {
		resp = response.BuildVideoInfoUpdateResp(err)
		return resp, err
	}
	resp = response.BuildVideoInfoUpdateResp(nil)
	return
}

// UpdateVideoCommentCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) UpdateVideoCommentCount(ctx context.Context, req *VideoService.UpdateVideoCommentCountRequest) (resp *VideoService.UpdateVideoInfoResponse, err error) {
	// TODO: Your code here...
	err = service.NewUpdateVideoInfoService(ctx).UpdateVideoCommentCount(req)
	if err != nil {
		resp = response.BuildVideoInfoUpdateResp(err)
		return resp, err
	}
	resp = response.BuildVideoInfoUpdateResp(nil)
	return
}

// GetAuthorIdByVideoId implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetAuthorIdByVideoId(ctx context.Context, req *VideoService.GetAuthorIdByVideoIdRequest) (resp *VideoService.GetAuthorIdByVideoIdResponse, err error) {
	// TODO: Your code here...

	if req.VideoId <= 0 {
		resp = response.BuildGetAuthorIdByVideoIdResponse(-1, errno.ErrInvalidParameter)
		return resp, err
	}

	authorId, err := service.NewGetVideoInfoService(ctx).GetAuthorID(req)
	if err != nil {
		resp = response.BuildGetAuthorIdByVideoIdResponse(-1, err)
		return resp, err
	}

	resp = response.BuildGetAuthorIdByVideoIdResponse(authorId, nil)

	return resp, nil
}
