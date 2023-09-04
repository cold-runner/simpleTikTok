package request

import (
	"github.com/cold-runner/simpleTikTok/kitex_gen/VideoService"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/social/dal"
	"github.com/cold-runner/simpleTikTok/service/social/rpc"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v3"
)

type CommentRequestToDTM struct {
	CommentText string `json:"comment_text"`
	VideoId     int64  `json:"video_id"`
	UserId      int64  `json:"user_id"`
	CommentId   int64  `json:"comment_id"`
	ActionType  int32  `json:"action_type"`
}

var (
	// 1. 在评论表中添加评论记录
	addComment    = "/addComment"
	deleteComment = "/deleteComment"
	// 2. 在视频表中的评论数 comment_count +1 (rpc 调用)
	addVideoCommentCount    = "/addVideoCommentCount"
	deleteVideoCommentCount = "/deleteVideoCommentCount"
)

func CommentActionRequest(request *CommentRequestToDTM) error {
	var saga *dtmcli.Saga
	gid := shortuuid.New()
	actionType := request.ActionType
	req := &gin.H{
		"user_id":      request.UserId,
		"video_id":     request.VideoId,
		"comment_text": request.CommentText,
		"comment_id":   request.CommentId,
	}
	log.Debugw("start comment action request, action type is", "action_type", request.ActionType)
	if actionType == 1 {
		saga = dtmcli.NewSaga(dtmServer, gid).
			Add(host+addComment,
				host+abortedCompensate,
				req).
			Add(host+addVideoCommentCount,
				host+deleteComment,
				req)
	} else if actionType == 2 {
		saga = dtmcli.NewSaga(dtmServer, gid).
			Add(host+deleteComment,
				host+abortedCompensate,
				req).
			Add(host+deleteVideoCommentCount,
				host+addComment,
				req)
	}
	saga.WaitResult = true
	err := saga.Submit()
	if err != nil {
		log.Errorw("submit comment action request failed", "err", err)
		return err
	}
	return nil
}

func AddCommentActionRoute(app *gin.Engine) {
	// 1. 在评论表中添加评论记录
	app.POST(api+addComment, func(c *gin.Context) {
		var userId int64
		var videoId int64
		var commentText string
		if _, exists := c.Get("userID"); exists {
			// 补偿操作
			log.Debugw("compensate: start add comment  request")
			userId = c.GetInt64("userID")
			videoId = c.GetInt64("videoID")
			content := c.GetString("commentContent")
			_, err := dal.CreateComment(c, videoId, userId, content)
			if err != nil {
				log.Errorw("compensate: add comment failed",
					"err", err)
				c.JSON(409, err.Error())
			} else {
				c.JSON(200, "compensate: add comment success")
			}
		} else {
			log.Debugw("start add comment action request")
			var data map[string]interface{}
			if err := c.ShouldBindJSON(&data); err != nil {
				log.Errorw("Failed to bind JSON: %v", err)
			}
			userId = int64(data["user_id"].(float64))
			videoId = int64(data["video_id"].(float64))
			commentText = data["comment_text"].(string)
			commentID, err := dal.CreateComment(c, videoId, userId, commentText)
			if err != nil {
				log.Errorw("add comment failed", "err", err)
				c.JSON(409, err.Error())
			} else {
				c.Set("commentID", commentID)
				log.Debugw("add comment success")
				c.JSON(200, "add comment success")
			}
		}
	})
	// 1. 在评论表中删除评论记录
	app.POST(api+deleteComment, func(c *gin.Context) {
		if _, exists := c.Get("commentID"); exists {
			log.Debugw("compensate: start delete comment action request")
			commentId := c.GetInt64("commentID")
			_, err := dal.DeleteComment(c, commentId)
			if err != nil {
				log.Errorw("compensate: delete comment failed", "err", err)
				c.JSON(409, err.Error())
			} else {
				c.JSON(200, "compensate: delete comment success")
			}
		} else {
			log.Debugw("start delete comment action request")
			var data map[string]interface{}
			if err := c.ShouldBindJSON(&data); err != nil {
				log.Errorw("Failed to bind JSON: %v", err)
			}
			commentId := int64(data["comment_id"].(float64))
			comment, err := dal.DeleteComment(c, commentId)
			if err != nil {
				log.Errorw("delete comment failed", "err", err)
				c.JSON(409, err.Error())
			} else {
				c.Set("userID", comment.AuthorId)
				c.Set("videoID", comment.VideoId)
				c.Set("commentContent", comment.Content)
				log.Debugw("delete comment success")
				c.JSON(200, "delete comment success")
			}
		}
	})

	// 2. 更新视频表的评论数
	app.POST(api+addVideoCommentCount, func(c *gin.Context) {
		log.Debugw("start add video comment count")
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			log.Errorw("Failed to bind JSON: %v", err)
		}
		videoId := int64(data["video_id"].(float64))
		_, err := rpc.VideoClient.UpdateVideoCommentCount(c, &VideoService.UpdateVideoCommentCountRequest{
			VideoId:    videoId,
			ActionType: 1,
		})
		if err != nil {
			log.Errorw("update video comment count failed", "err", err)
			c.JSON(409, err.Error())
		} else {
			log.Debugw("add video comment count success")
			c.JSON(200, "add video comment count success")
		}
	})

	app.POST(api+deleteVideoCommentCount, func(c *gin.Context) {
		log.Debugw("start delete video comment count")
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			log.Errorw("Failed to bind JSON: %v", err)
		}
		videoId := int64(data["video_id"].(float64))
		_, err := rpc.VideoClient.UpdateVideoCommentCount(c, &VideoService.UpdateVideoCommentCountRequest{
			VideoId:    videoId,
			ActionType: 2,
		})
		if err != nil {
			log.Errorw("update video comment count failed", "err", err)
			c.JSON(409, err.Error())
		} else {
			log.Debugw("delete video comment count success")
			c.JSON(200, "delete video comment count success")
		}
	})
}
