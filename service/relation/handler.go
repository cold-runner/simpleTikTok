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
