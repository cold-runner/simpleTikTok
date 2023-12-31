// Code generated by hertz generator.

package ApiServer

import (
	"context"
	"github.com/cold-runner/simpleTikTok/apiServer/biz/handler/response"
	"github.com/cold-runner/simpleTikTok/apiServer/rpc"
	"github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	mw "github.com/cold-runner/simpleTikTok/pkg/middleware"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cold-runner/simpleTikTok/apiServer/biz/model/ApiServer"
)

// FavoriteAction .
// @router /douyin/favorite/action/ [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ApiServer.FavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	v, _ := c.Get(mw.IdentityKey)
	_, err = rpc.FavoriteAction(context.Background(),
		&SocialService.FavoriteActionRequest{
			UserId:     v.(*ApiServer.User).Id,
			ToVideoId:  req.VideoId,
			ActionType: req.ActionType,
		})

	if err != nil {
		log.Errorw("rpc favorite action failed", "err", err)
		response.SendFavoriteActionResponse(c, err)
		return
	}
	response.SendFavoriteActionResponse(c, errno.OK)
}

// FavoriteList .
// @router /douyin/favorite/list/ [GET]
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ApiServer.FavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	v, _ := c.Get(mw.IdentityKey)
	listResp, err := rpc.FavoriteList(context.Background(),
		&SocialService.FavoriteListRequest{
			UserId:   v.(*ApiServer.User).Id,
			ToUserId: req.GetUserId(),
		})
	if err != nil {
		log.Errorw("rpc favorite list failed", "err", err)
		response.SendFavoriteListResponse(c, nil, err)
	} else {
		response.SendFavoriteListResponse(c, listResp, errno.OK)
	}

}

// CommentAction .
// @router /douyin/comment/action/ [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ApiServer.CommentActionRequest
	actionTypeStr := c.Query("action_type")
	actionType, _ := strconv.ParseInt(actionTypeStr, 10, 64)

	if actionType == 1 {
		videoIdStr := c.Query("video_id")
		videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)
		req.VideoId = videoId
		req.CommentText = c.Query("comment_text")

	} else if actionType == 2 {
		commentIdStr := c.Query("comment_id")
		commentId, _ := strconv.ParseInt(commentIdStr, 10, 64)
		req.CommentId = commentId
		videoIdStr := c.Query("video_id")
		videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)
		req.VideoId = videoId
	} else {
		c.String(consts.StatusBadRequest, "action type error")
		return
	}
	v, _ := c.Get(mw.IdentityKey)
	resp, err := rpc.CommentAction(context.Background(),
		&SocialService.CommentActionRequest{
			UserId:      v.(*ApiServer.User).Id,
			VideoId:     req.GetVideoId(),
			ActionType:  int32(actionType),
			CommentText: req.GetCommentText(),
			CommentId:   req.GetCommentId(),
		})
	log.Debugw("resp from rpc comment action", "resp", resp)

	if err != nil {
		log.Errorw("rpc comment action failed", "err", err)
		response.SendCommentActionResponse(c, resp, err)
	} else {
		response.SendCommentActionResponse(c, resp, errno.OK)
	}
}

// VideoCommentList .
// @router /douyin/comment/list/ [GET]
func VideoCommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ApiServer.VideoCommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	v, _ := c.Get(mw.IdentityKey)
	resp, err := rpc.CommentList(context.Background(), &SocialService.CommentListRequest{
		UserId:  v.(*ApiServer.User).Id,
		VideoId: req.GetVideoId(),
	})
	log.Debugw("resp from rpc comment list", "resp", resp)

	if err != nil {
		log.Errorw("rpc comment list failed", "err", err)
		response.SendCommentListResponse(c, resp, err)
	} else {
		response.SendCommentListResponse(c, resp, errno.OK)
	}

}
