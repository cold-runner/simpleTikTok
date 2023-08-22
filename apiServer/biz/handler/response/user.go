package responseUtil

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
)

// 注册用户响应报文
type RegisterResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
}

// 发送注册响应报文
func SendRegisterResponse(c *app.RequestContext, error error, id int64,
	token string) {
	err := errno.MatchErr(error)
	c.JSON(200, RegisterResponse{
		StatusCode: err.HTTP,
		StatusMsg:  err.Message,
		UserId:     id,
		Token:      token,
	})
}
