package pack

import (
	"github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
	"github.com/cold-runner/simpleTikTok/pkg/dao/model"
)

func CommentInfoConvert(user *SocialService.User, v *model.Comment) *SocialService.Comment {
	return &SocialService.Comment{
		Id:         v.ID,
		User:       user,
		Content:    v.Content,
		CreateDate: v.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
