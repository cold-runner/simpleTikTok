package social

import (
	"context"
	SocialService "github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
)

// SocialServiceImpl implements the last service interface defined in the IDL.
type SocialServiceImpl struct{}

// CommentAction implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) CommentAction(ctx context.Context, req *SocialService.CommentActionRequest) (resp *SocialService.CommentActionResponse, err error) {

	return
}

// CommentList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) CommentList(ctx context.Context, req *SocialService.CommentListRequest) (resp *SocialService.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}

// LikeAction implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) LikeAction(ctx context.Context, req *SocialService.FavoriteActionRequest) (resp *SocialService.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
}

// LikeList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) LikeList(ctx context.Context, req *SocialService.FavoriteListRequest) (resp *SocialService.FavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}
