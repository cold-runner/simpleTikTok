package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/VideoService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
)

type FavoriteActionResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}
type FavoriteListResponse struct {
	StatusCode int32                 `json:"status_code"`
	StatusMsg  string                `json:"status_msg"`
	VideoList  []*VideoService.Video `json:"video_list"`
}

type CommentActionResponse struct {
	StatusCode int32                  `json:"status_code"`
	StatusMsg  string                 `json:"status_msg"`
	Comment    *SocialService.Comment `json:"comment"`
}

type CommentListResponse struct {
	StatusCode  int32                    `json:"status_code"`
	StatusMsg   string                   `json:"status_msg"`
	CommentList []*SocialService.Comment `json:"comment_list"`
}

func SendFavoriteActionResponse(c *app.RequestContext, err error) {
	Err := errno.MatchErr(err)
	c.JSON(consts.StatusOK, FavoriteActionResponse{
		StatusCode: Err.HTTP,
		StatusMsg:  Err.Message,
	})
}

func SendFavoriteListResponse(c *app.RequestContext, resp *SocialService.FavoriteListResponse, err error) {
	Err := errno.MatchErr(err)
	if Err.HTTP == 0 || Err.HTTP == 200 {
		c.JSON(consts.StatusOK, VideoPublishListResponse{
			StatusCode: 0,
			StatusMsg:  "Get Favorite List Success!",
			VideoList:  resp.GetVideoList(),
		})

	} else {
		c.JSON(int(Err.HTTP), VideoPublishListResponse{
			StatusCode: Err.HTTP,
			StatusMsg:  Err.Message,
			VideoList:  nil,
		})
	}
}

func SendCommentActionResponse(c *app.RequestContext, resp *SocialService.CommentActionResposne, err error) {
	Err := errno.MatchErr(err)
	if (Err.HTTP == 0 || Err.HTTP == 200) && resp.Comment == nil {
		c.JSON(consts.StatusOK, CommentActionResponse{
			StatusCode: 0,
			StatusMsg:  "delete Comment Success!",
			Comment:    nil,
		})

	} else if (Err.HTTP == 0 || Err.HTTP == 200) && resp.Comment != nil {
		c.JSON(consts.StatusOK, CommentActionResponse{
			StatusCode: Err.HTTP,
			StatusMsg:  Err.Message,
			Comment:    resp.GetComment(),
		})
	} else {
		c.JSON(int(Err.HTTP), CommentActionResponse{
			StatusCode: Err.HTTP,
			StatusMsg:  Err.Message,
			Comment:    nil,
		})
	}
}

func SendCommentListResponse(c *app.RequestContext,
	resp *SocialService.CommentListResponse, err error) {
	Err := errno.MatchErr(err)

	if (Err.HTTP == 0 || Err.HTTP == 200) && resp.CommentList != nil {
		c.JSON(consts.StatusOK, CommentListResponse{
			StatusCode:  0,
			StatusMsg:   "Get Comment List Success!",
			CommentList: resp.GetCommentList(),
		})
	} else if (Err.HTTP == 0 || Err.HTTP == 200) && resp.CommentList != nil {
		c.JSON(consts.StatusOK, CommentListResponse{
			StatusCode:  0,
			StatusMsg:   "Get Comment List Success! But no comment",
			CommentList: nil,
		})
	} else {
		c.JSON(int(Err.HTTP), CommentListResponse{
			StatusCode:  Err.HTTP,
			StatusMsg:   Err.Message,
			CommentList: nil,
		})
	}
}
