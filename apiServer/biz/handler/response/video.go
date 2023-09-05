package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cold-runner/simpleTikTok/kitex_gen/VideoService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
)

type PublishActionResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type FeedResponse struct {
	StatusCode int32                 `json:"status_code"`
	StatusMsg  string                `json:"status_msg"`
	NextTime   int64                 `json:"next_time"`
	VideoList  []*VideoService.Video `json:"video_list"`
}

type VideoPublishListResponse struct {
	StatusCode int32                 `json:"status_code"`
	StatusMsg  string                `json:"status_msg"`
	VideoList  []*VideoService.Video `json:"video_list"`
}

func SendPublishActionResponse(c *app.RequestContext, err error) {
	Err := errno.MatchErr(err)
	c.JSON(consts.StatusOK, PublishActionResponse{
		StatusCode: Err.HTTP,
		StatusMsg:  Err.Message,
	})
}

func SendFeedResponse(c *app.RequestContext,
	videoList []*VideoService.Video, nextTime int64, err error) {
	Err := errno.MatchErr(err)
	if len(videoList) == 0 {
		c.JSON(consts.StatusOK, FeedResponse{
			StatusCode: Err.HTTP,
			StatusMsg:  "Get Feed Success!",
			NextTime:   nextTime,
			VideoList:  nil,
		})
		return
	}
	c.JSON(consts.StatusOK, FeedResponse{
		StatusCode: Err.HTTP,
		StatusMsg:  "Get Feed Success!",
		NextTime:   nextTime,
		VideoList:  videoList,
	})
}

func SendVideoPublishListResponse(c *app.RequestContext,
	videoList []*VideoService.Video, err error) {
	Err := errno.MatchErr(err)
	if len(videoList) == 0 {
		c.JSON(consts.StatusOK, FeedResponse{
			StatusCode: Err.HTTP,
			StatusMsg:  "Get Video List Success, but no video!",
			VideoList:  nil,
		})
		return
	}
	c.JSON(consts.StatusOK, FeedResponse{
		StatusCode: Err.HTTP,
		StatusMsg:  "Get Video List Success!",
		VideoList:  videoList,
	})
}
