package main

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/RelationService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/relation/response"
	"github.com/cold-runner/simpleTikTok/service/relation/service"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *RelationService.RelationActionRequest) (resp *RelationService.RelationActionResponse, err error) {
	// TODO: Your code here...
	uid, toUid, actionType := req.Uid, req.ToUserId, req.ActionType
	if uid <= 0 || toUid <= 0 {
		log.Errorw("RelationActionService.RelationAction", "err", errno.ErrInvalidParameter)
		return nil, errno.ErrInvalidParameter
	}
	if actionType <= 0 || actionType > 2 {
		log.Errorw("RelationActionService.RelationAction", "err", errno.ErrInvalidParameter)
		return nil, errno.ErrInvalidParameter
	}

	if uid == toUid {
		log.Errorw("RelationActionService.RelationAction", "err", errno.ErrInvalidParameter)
		return nil, errno.ErrInvalidParameter
	}

	err = service.NewRelationActionService(ctx).RelationAction(req)
	if err != nil {
		resp = response.BuildRelationActionResp(err)
		return resp, err
	}
	resp = response.BuildRelationActionResp(errno.OK)
	return resp, nil

}

// GetFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowList(ctx context.Context, req *RelationService.RelationFollowListRequest) (resp *RelationService.RelationFollowListResponse, err error) {
	// TODO: Your code here...
	userInfos, err := service.NewGetFollowListService(ctx).GetUserFollowList(req)
	if err != nil {
		resp = response.BuildGetUserFollowListResp(nil, err)
		return resp, nil
	}
	resp = response.BuildGetUserFollowListResp(userInfos, errno.OK)
	return resp, nil
}

// GetFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowerList(ctx context.Context, req *RelationService.RelationFollowerListRequest) (resp *RelationService.RelationFollowerListResponse, err error) {
	// TODO: Your code here...
	userInfos, err := service.NewGetFollowerListService(ctx).GetUserFollowerList(req)
	if err != nil {
		resp = response.BuildGetUserFollowerListResp(nil, err)
		return resp, nil
	}
	resp = response.BuildGetUserFollowerListResp(userInfos, errno.OK)
	return resp, nil
}

// QueryRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) QueryRelation(ctx context.Context, req *RelationService.RelationQueryRequest) (resp *RelationService.RelationQueryResponse, err error) {
	// TODO: Your code here...
	uid, toUid := req.Uid, req.ToUserId
	if uid <= 0 || toUid <= 0 {
		log.Errorw("RelationActionService.QueryRelation", "err", errno.ErrInvalidParameter)
		return nil, errno.ErrInvalidParameter
	}

	isFollow, err := service.NewQueryRelationService(ctx).QueryRelation(req)
	if err != nil {
		resp = response.BuildQueryActionResp(false, err)
		return resp, err
	}
	resp = response.BuildQueryActionResp(isFollow, errno.OK)
	return resp, nil
}

// QueryUserInfosWithRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) QueryUserInfosWithRelation(ctx context.Context, req *RelationService.QueryUserInfosWithRelationRequest) (resp *RelationService.QueryUserInfosWithRelationResponse, err error) {
	// TODO: Your code here...
	uid, toUids := req.GetUid(), req.GetIdList()
	if uid <= 0 || len(toUids) <= 0 {
		log.Errorw("RelationActionService.QueryUserInfosWithRelation",
			"err", errno.ErrInvalidParameter, "uid:", uid, "toUids:", toUids)
		return nil, errno.ErrInvalidParameter
	}
	uInfos, err := service.NewGetUserInfoWithRelationService(ctx).GetUserInfoList(req)
	if err != nil {
		log.Errorw("RelationActionService.QueryUserInfosWithRelation", "err", err)
		resp = response.BuildGetUserInfosWithRelationResp(nil, err)
		return resp, err
	}
	log.Debugw("Get user info with relation success", "info number:", len(uInfos))
	resp = response.BuildGetUserInfosWithRelationResp(uInfos, nil)
	return resp, nil
}
