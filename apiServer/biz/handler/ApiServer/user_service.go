package ApiServer

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cold-runner/simpleTikTok/apiServer/biz/handler/response"
	"github.com/cold-runner/simpleTikTok/apiServer/biz/model/ApiServer"
	"github.com/cold-runner/simpleTikTok/apiServer/rpc"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	mw "github.com/cold-runner/simpleTikTok/pkg/middleware"
)

// Register .
// @router /douyin/user/register/ [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ApiServer.UserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	log.Infow("Calling remote rpc register request:", "req", req)
	userId, err := rpc.Register(context.Background(),
		&UserService.UserRegisterRequest{
			Username: req.Username,
			Password: req.Password,
		})
	if err != nil {
		response.SendRegisterResponse(c, errno.MatchErr(err), -1, "")
		return
	}
	// 签发jwt 验证 token
	log.Infow(" Generating token for user.")
	token, _, err := mw.JwtMiddleware.TokenGenerator(userId)
	if err != nil {
		response.SendRegisterResponse(c, errno.MatchErr(err), -1, "")
		return
	}

	response.SendRegisterResponse(c, nil, userId, token)
}

// Login .
// @router /douyin/user/login/ [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	mw.JwtMiddleware.LoginHandler(ctx, c)
}

// GetUserInfo .
// @router /douyin/user/ [GET]
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ApiServer.UserInfoRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	v, _ := c.Get(mw.IdentityKey)
	log.Infow("Calling remote rpc get user info request:", "req", req)
	resp, err := rpc.GetUserInfo(context.Background(),
		&UserService.UserInfoRequest{
			FromId: v.(*ApiServer.User).Id,
			ToId:   req.UserId,
		})

	if err != nil {
		response.SendUserInfoResponse(c, err, nil)
		return
	}
	response.SendUserInfoResponse(c, nil, resp)
}
