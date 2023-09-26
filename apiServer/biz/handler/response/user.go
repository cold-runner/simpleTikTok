package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
)

// RegisterResponse 注册用户响应报文
type RegisterResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
}

type userInfo struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

// UserInfoResponse 获取用户信息响应报文
type UserInfoResponse struct {
	StatusCode int32    `json:"status_code"`
	StatusMsg  string   `json:"status_msg"`
	User       userInfo `json:"user"`
}

// SendRegisterResponse 发送注册响应报文
func SendRegisterResponse(c *app.RequestContext, error error, id int64,
	token string) {
	if error != nil {
		err := errno.MatchErr(error)
		c.JSON(int(err.HTTP), RegisterResponse{
			StatusCode: err.HTTP,
			StatusMsg:  err.Message,
			UserId:     id,
			Token:      token,
		})
	} else {
		c.JSON(200, RegisterResponse{
			StatusCode: 0,
			StatusMsg:  "User registration successfully! ",
			UserId:     id,
			Token:      token,
		})
	}
}

func SendUserInfoResponse(c *app.RequestContext, error error,
	resp *UserService.UserInfoResponse) {
	if error != nil {
		err := errno.MatchErr(error)
		c.JSON(int(err.HTTP), UserInfoResponse{
			StatusCode: err.HTTP,
			StatusMsg:  err.Message,
			User:       userInfo{},
		})
	} else {
		c.JSON(200, UserInfoResponse{
			StatusCode: 0,
			StatusMsg:  "Get user info successfully! ",
			User: userInfo{
				ID:              resp.User.Id,
				Name:            resp.User.Name,
				FollowCount:     resp.User.FollowCount,
				FollowerCount:   resp.User.FollowerCount,
				IsFollow:        resp.User.IsFollow,
				Avatar:          resp.User.Avatar,
				BackgroundImage: resp.User.BackgroundImage,
				Signature:       resp.User.Signature,
				TotalFavorited:  resp.User.TotalFavorited,
				WorkCount:       resp.User.WorkCount,
				FavoriteCount:   resp.User.FavoriteCount,
			},
		})
	}
}
