package main

import (
	"context"
	ApiServer "github.com/cold-runner/simpleTikTok/kitex_gen/ApiServer"
	SocialService "github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
	VideoService "github.com/cold-runner/simpleTikTok/kitex_gen/VideoService"
	RelationService "github.com/cold-runner/simpleTikTok/kitex_gen/RelationService"
	UserService "github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *UserService.UserRegisterRequest) (resp *UserService.UserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *UserService.UserLoginRequest) (resp *UserService.UserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *UserService.UserInfoRequest) (resp *UserService.UserInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// MGetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUserInfo(ctx context.Context, req *UserService.MGetUserInfoRequest) (resp *UserService.MGetUserInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// ChangeUserFollowCount implements the UserServiceImpl interface.
func (s *UserServiceImpl) ChangeUserFollowCount(ctx context.Context, req *UserService.ChangeUserFollowCountRequest) (resp *ApiServer.BaseResp, err error) {
	// TODO: Your code here...
	return
}

// ChangeUserTotalFavoritedCount implements the UserServiceImpl interface.
func (s *UserServiceImpl) ChangeUserTotalFavoritedCount(ctx context.Context, req *UserService.ChangeUserTotalFavoritedCountRequest) (resp *UserService.UpdateUserInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// ChangeUserWorkCount implements the UserServiceImpl interface.
func (s *UserServiceImpl) ChangeUserWorkCount(ctx context.Context, req *UserService.ChangeUserWorkCountRequest) (resp *UserService.UpdateUserInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// ChangeUserFavoriteCount implements the UserServiceImpl interface.
func (s *UserServiceImpl) ChangeUserFavoriteCount(ctx context.Context, req *UserService.ChangeUserFavoriteCountRequest) (resp *UserService.UpdateUserInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *RelationService.RelationActionRequest) (resp *RelationService.RelationActionResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowList(ctx context.Context, req *RelationService.RelationFollowListRequest) (resp *RelationService.RelationFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowerList(ctx context.Context, req *RelationService.RelationFollowerListRequest) (resp *RelationService.RelationFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) QueryRelation(ctx context.Context, req *RelationService.RelationQueryRequest) (resp *RelationService.RelationQueryResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryUserInfosWithRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) QueryUserInfosWithRelation(ctx context.Context, req *RelationService.QueryUserInfosWithRelationRequest) (resp *RelationService.QueryUserInfosWithRelationResponse, err error) {
	// TODO: Your code here...
	return
}

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
