package dtmRequest

import (
	"fmt"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/VideoService"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/social/dal"
	"github.com/cold-runner/simpleTikTok/service/social/rpc"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v3"
	"github.com/spf13/viper"
)

type FavoriteActionRequestToDTM struct {
	UserId     int64 `json:"user_id"`
	VideoId    int64 `json:"video_id"`
	AuthorId   int64 `json:"author_id"`
	ActionType int32 `json:"action_type"`
}

// 给一个视频点赞需要完成如下的业务逻辑：
var (
	// 1. 在视频点赞关系表中添加一条记录 （mysql 数据库操作)
	addVideoFavoriteRelation    = "/addVideoFavoriteRelation"
	deleteVideoFavoriteRelation = "/deleteVideoFavoriteRelation"
	// 2. 在视频表中的点赞数 favorite_count +1 (rpc 调用)
	addVideoFavoriteCount    = "/addVideoFavoriteCount"
	deleteVideoFavoriteCount = "/deleteVideoFavoriteCount"
	// 3. 在用户表中的用户喜欢数favorite_count +1 (rpc 调用)
	addUserFavoriteCount    = "/addUserFavoriteCount"
	deleteUserFavoriteCount = "/deleteUserFavoriteCount"
	// 4. 在用户表中的作者获赞数量 total_favorited +1 (rpc 调用)
	addAuthorTotalFavorited    = "/addAuthorTotalFavorited"
	deleteAuthorTotalFavorited = "/deleteAuthorTotalFavorited"
)

func FavoriteActionRequest(request *FavoriteActionRequestToDTM) error {
	var saga *dtmcli.Saga
	gid := shortuuid.New()
	req := &gin.H{
		"user_id":     request.UserId,
		"video_id":    request.VideoId,
		"author_id":   request.AuthorId,
		"action_type": request.ActionType,
	}
	dtmServer := viper.GetString("dtm.server")
	host := fmt.Sprintf(
		"%s:%d%s",
		viper.GetString("dtm.http.host"),
		viper.GetInt("dtm.http.port"),
		viper.GetString("dtm.http.api"),
	)

	log.Debugw("host address:", "host", host)
	//req := &gin.H{"amount": 1}
	if request.ActionType == 1 {
		saga = dtmcli.NewSaga(dtmServer, gid).
			Add(host+addVideoFavoriteRelation,
				host+deleteVideoFavoriteRelation,
				req).
			Add(host+addVideoFavoriteCount,
				host+deleteVideoFavoriteCount,
				req).
			Add(host+addUserFavoriteCount,
				host+deleteUserFavoriteCount,
				req).
			Add(host+addAuthorTotalFavorited,
				host+deleteAuthorTotalFavorited,
				req)
	} else if request.ActionType == 2 {
		saga = dtmcli.NewSaga(dtmServer, gid).
			Add(host+deleteVideoFavoriteRelation,
				host+addVideoFavoriteRelation,
				req).
			Add(host+deleteVideoFavoriteCount,
				host+addVideoFavoriteCount,
				req).
			Add(host+deleteUserFavoriteCount,
				host+addUserFavoriteCount,
				req).
			Add(host+deleteAuthorTotalFavorited,
				host+addAuthorTotalFavorited,
				req)
	}
	saga.WaitResult = true
	err := saga.Submit()
	if err != nil {
		log.Errorw("submit favorite saga failed", "err", err)
		return err
	}
	return nil
}

func AddFavoriteActionRoute(app *gin.Engine) {
	// 1. 视频点赞关系表
	api := viper.GetString("dtm.http.api")

	app.POST(api+addVideoFavoriteRelation, func(c *gin.Context) {
		log.Debugw("start add video favorite relation")
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			log.Errorw("Failed to bind JSON: %v", err)
		}
		userId := int64(data["user_id"].(float64))
		videoId := int64(data["video_id"].(float64))

		err := dal.AddFavorite(c, userId, videoId)
		if err != nil {
			log.Errorw("add favorite failed", "err", err)
			c.JSON(409, err.Error())
		}
		log.Debugw("add favorite relation success")
		c.JSON(200, "add video favorite relation success")
	})

	app.POST(api+deleteVideoFavoriteRelation, func(c *gin.Context) {
		log.Debugw("start delete video favorite relation")
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			log.Errorw("Failed to bind JSON: %v", err)
		}
		userId := int64(data["user_id"].(float64))
		videoId := int64(data["video_id"].(float64))

		err := dal.RemoveFavorite(c, userId, videoId)
		if err != nil {
			log.Errorw("add favorite failed", "err", err)
			c.JSON(409, err.Error())
		}
		log.Debugw("delete favorite relation success")
		c.JSON(200, "delete video favorite relation success")
	})

	// 2. 视频表
	app.POST(api+addVideoFavoriteCount, func(c *gin.Context) {
		log.Debugw("start add video favorite count")
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			log.Errorw("Failed to bind JSON: %v", err)
		}
		videoId := int64(data["video_id"].(float64))
		resp, err := rpc.VideoClient.UpdateVideoFavoriteCount(c,
			&VideoService.UpdateVideoFavoriteCountRequest{
				VideoId:    videoId,
				ActionType: 1,
			})
		if err != nil {
			log.Errorw("update video favorite count failed", "err", err)
			c.JSON(409, err)
		}
		if resp.BaseResp.GetStatusCode() != 0 || resp.BaseResp.
			GetStatusCode() != 200 {
			c.JSON(409, resp)
		}
		log.Debugw("add video favorite count success")
		c.JSON(200, "update video favorite count success")
	})
	app.POST(api+deleteVideoFavoriteCount, func(c *gin.Context) {
		log.Debugw("start delete video favorite count")
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			log.Errorw("Failed to bind JSON: %v", err)
		}
		videoId := int64(data["video_id"].(float64))
		resp, err := rpc.VideoClient.UpdateVideoFavoriteCount(c,
			&VideoService.UpdateVideoFavoriteCountRequest{
				VideoId:    videoId,
				ActionType: 2,
			})
		if err != nil {
			log.Errorw("update video favorite count failed", "err", err)
			c.JSON(409, err)
		}
		if resp.BaseResp.GetStatusCode() != 0 || resp.BaseResp.
			GetStatusCode() != 200 {
			c.JSON(409, resp)
		}
		log.Debugw("delete video favorite count success")
		c.JSON(200, "update video favorite count success")
	})
	// 3. 用户表用户
	app.POST(api+addUserFavoriteCount, func(c *gin.Context) {
		log.Debugw("start add user favorite count")
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			log.Errorw("Failed to bind JSON: %v", err)
		}
		userId := int64(data["user_id"].(float64))
		resp, err := rpc.UserClient.ChangeUserFavoriteCount(c,
			&UserService.ChangeUserFavoriteCountRequest{
				Id:         userId,
				ActionType: 1,
			})
		if err != nil {
			log.Errorw("update user favorite count failed", "err", err)
			c.JSON(409, err.Error())
		}
		if resp.BaseResp.GetStatusCode() != 0 || resp.BaseResp.
			GetStatusCode() != 200 {
			c.JSON(409, resp)
		}
		log.Debugw("add user favorite count success")
		c.JSON(200, "update user favorite count success")
	})

	app.POST(api+deleteUserFavoriteCount, func(c *gin.Context) {
		log.Debugw("start delete user favorite count")
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			log.Errorw("Failed to bind JSON: %v", err)
		}
		userId := int64(data["user_id"].(float64))
		resp, err := rpc.UserClient.ChangeUserFavoriteCount(c,
			&UserService.ChangeUserFavoriteCountRequest{
				Id:         userId,
				ActionType: 2,
			})
		if err != nil {
			log.Errorw("update user favorite count failed", "err", err)
			c.JSON(409, err.Error())
		}
		if resp.BaseResp.GetStatusCode() != 0 || resp.BaseResp.
			GetStatusCode() != 200 {
			c.JSON(409, resp)
		}
		log.Debugw("delete user favorite count success")
		c.JSON(200, "update user favorite count success")
	})

	// 4. 用户表作者
	app.POST(api+addAuthorTotalFavorited, func(c *gin.Context) {
		log.Debugw("start add author total favorited")
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			log.Errorw("Failed to bind JSON: %v", err)
		}
		authorId := int64(data["author_id"].(float64))
		resp, err := rpc.UserClient.ChangeUserTotalFavoritedCount(c,
			&UserService.ChangeUserTotalFavoritedCountRequest{
				Id:         authorId,
				ActionType: 1,
			})
		if err != nil {
			log.Errorw("update user favorite count failed", "err", err)
			c.JSON(409, err.Error())
		}
		if resp.BaseResp.GetStatusCode() != 0 || resp.BaseResp.
			GetStatusCode() != 200 {
			c.JSON(409, resp)
		}
		log.Debugw("add author total favorited success")
		c.JSON(200, "update user favorite count success")
	})

	app.POST(api+deleteAuthorTotalFavorited, func(c *gin.Context) {
		log.Debugw("start add author total favorited")
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			log.Errorw("Failed to bind JSON: %v", err)
		}
		authorId := int64(data["author_id"].(float64))
		resp, err := rpc.UserClient.ChangeUserTotalFavoritedCount(c,
			&UserService.ChangeUserTotalFavoritedCountRequest{
				Id:         authorId,
				ActionType: 2,
			})
		if err != nil {
			log.Errorw("update user favorite count failed", "err", err)
			c.JSON(409, err.Error())
		}
		if resp.BaseResp.GetStatusCode() != 0 || resp.BaseResp.
			GetStatusCode() != 200 {
			c.JSON(409, resp)
		}
		log.Debugw("add author total favorited success")
		c.JSON(200, "update user favorite count success")
	})
}
