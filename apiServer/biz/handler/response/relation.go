package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cold-runner/simpleTikTok/kitex_gen/BaseResponse"
	"github.com/cold-runner/simpleTikTok/kitex_gen/RelationService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/pkg/log"
)

type RelationActionResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type FollowListResponse struct {
	StatusCode int32               `json:"status_code"`
	StatusMsg  string              `json:"status_msg"`
	UserList   []*UserService.User `json:"user_list"`
}
type FollowerListResponse struct {
	StatusCode int32               `json:"status_code"`
	StatusMsg  string              `json:"status_msg"`
	UserList   []*UserService.User `json:"user_list"`
}

func SendRelationActionResponse(c *app.RequestContext, resp interface{}) {
	switch v := resp.(type) {
	case error:
		err := errno.MatchErr(v)
		c.JSON(int(err.HTTP), RelationActionResponse{
			StatusCode: err.HTTP,
			StatusMsg:  err.Message,
		})
	case *RelationService.RelationActionResponse:
		c.JSON(200, RelationActionResponse{
			StatusCode: 0,
			StatusMsg:  "Action Success!",
		})
	default:
		log.Errorw("Unknown type of response", "type", v)
		c.JSON(400, RelationActionResponse{
			StatusCode: -1,
			StatusMsg:  "Unknown type of response from rpc server!",
		})
	}
}

func SendFollowListResponse(c *app.RequestContext,
	resp *RelationService.RelationFollowListResponse, err error) {
	if err != nil {
		log.Errorw("api server rpc send follow list response failed", "err",
			err)
		err := errno.MatchErr(err)
		c.JSON(int(err.HTTP), RelationService.RelationFollowListResponse{
			BaseResp: &BaseResponse.BaseResp{
				StatusCode: err.HTTP,
				StatusMsg:  err.Message,
			},
			FollowList: nil,
		})
	} else if len(resp.FollowList) == 0 {
		log.Debugw("Sending a no follow list response", "resp", resp)
		// 处理用户没有关注的情况
		c.JSON(200, FollowListResponse{
			StatusCode: resp.BaseResp.StatusCode,
			StatusMsg:  "No follow yet!",
			UserList:   nil,
		})
	} else {
		log.Debugw("Sending a follow list response", "resp", resp)
		c.JSON(200, FollowListResponse{
			StatusCode: resp.BaseResp.GetStatusCode(),
			StatusMsg:  resp.BaseResp.GetStatusMsg(),
			UserList:   resp.FollowList,
		})
	}
}

func SendFollowerListResponse(c *app.RequestContext,
	resp *RelationService.RelationFollowerListResponse, err error) {
	if err != nil {
		log.Errorw("api server rpc send follower list response failed", "err", err)
		err := errno.MatchErr(err)
		c.JSON(int(err.HTTP), RelationService.RelationFollowerListResponse{
			BaseResp: &BaseResponse.BaseResp{
				StatusCode: err.HTTP,
				StatusMsg:  err.Message,
			},
			FollowerList: nil,
		})
	} else if len(resp.FollowerList) == 0 {
		log.Debugw("Sending a no follower list response", "resp", resp)
		c.JSON(200, FollowerListResponse{
			StatusCode: resp.BaseResp.GetStatusCode(),
			StatusMsg:  "No follower yet!",
			UserList:   nil,
		})
	} else {
		log.Debugw("Sending a follower list response", "resp", resp)
		c.JSON(200, FollowerListResponse{
			StatusCode: resp.BaseResp.GetStatusCode(),
			StatusMsg:  resp.BaseResp.GetStatusMsg(),
			UserList:   resp.FollowerList,
		})
	}
}
