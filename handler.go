package main

import (
	"context"
	VideoService "github.com/cold-runner/simpleTikTok/kitex_gen/VideoService"
	SocialService "github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// VideoFeed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoFeed(ctx context.Context, req *VideoService.VideoFeedRequest) (resp *VideoService.VideoFeedResponse, err error) {
	// TODO: Your code here...
	return
}

// VideoPublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoPublishAction(ctx context.Context, req *VideoService.VideoPublishActionRequest) (resp *VideoService.VideoPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// VideoPublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoPublishList(ctx context.Context, req *VideoService.VideoPublishListRequest) (resp *VideoService.VideoPublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// VideoPublishListByIds implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoPublishListByIds(ctx context.Context, req *VideoService.VideoPublishListByIdsRequest) (resp *VideoService.VideoPublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateVideoFavoriteCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) UpdateVideoFavoriteCount(ctx context.Context, req *VideoService.UpdateVideoFavoriteCountRequest) (resp *VideoService.UpdateVideoInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateVideoCommentCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) UpdateVideoCommentCount(ctx context.Context, req *VideoService.UpdateVideoCommentCountRequest) (resp *VideoService.UpdateVideoInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// GetAuthorIdByVideoId implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetAuthorIdByVideoId(ctx context.Context, req *VideoService.GetAuthorIdByVideoIdRequest) (resp *VideoService.GetAuthorIdByVideoIdResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteAction implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) FavoriteAction(ctx context.Context, req *SocialService.FavoriteActionRequest) (resp *SocialService.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentAction implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) CommentAction(ctx context.Context, req *SocialService.CommentActionRequest) (resp *SocialService.CommentActionResposne, err error) {
	// TODO: Your code here...
	return
}

// FavoriteList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) FavoriteList(ctx context.Context, req *SocialService.FavoriteListRequest) (resp *SocialService.FavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) CommentList(ctx context.Context, req *SocialService.CommentListRequest) (resp *SocialService.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}
