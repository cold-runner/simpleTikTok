package main

import (
	"context"
	SocialService "github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
)

// SocialServiceImpl implements the last service interface defined in the IDL.
type SocialServiceImpl struct{}

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
