package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
)

type FavoriteActionResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func SendFavoriteActionResponse(c *app.RequestContext, err error) {
	Err := errno.MatchErr(err)
	c.JSON(consts.StatusOK, FavoriteActionResponse{
		StatusCode: Err.HTTP,
		StatusMsg:  Err.Message,
	})
}
