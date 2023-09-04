package dal

import (
	"context"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/social/dal/model"
)

func CreateComment(ctx context.Context, videoId int64, authorId int64,
	content string) (int64, error) {
	comment := model.VideoComment{
		AuthorId: authorId,
		Content:  content,
		VideoId:  videoId,
	}
	// 创建一条新的评论记录
	if err := DB.WithContext(ctx).Create(&comment).Error; err != nil {
		log.Errorw("create comment failed", "err", err)
		return -1, err
	}
	return comment.CommentId, nil
}

func DeleteComment(ctx context.Context, commentId int64) (*model.VideoComment, error) {
	comment := &model.VideoComment{
		CommentId: commentId,
	}
	// 找到并填充评论信息
	if err := DB.WithContext(ctx).First(comment, commentId).Error; err != nil {
		log.Errorw("find comment failed", "err", err)
		return nil, err
	}

	// 删除评论记录
	if err := DB.WithContext(ctx).Delete(comment).Error; err != nil {
		log.Errorw("delete comment failed", "err", err)
		return nil, err
	}
	return comment, nil
}

func GetUserLatestComment(ctx context.Context,
	userId int64) (*model.VideoComment,
	error) {
	var comment model.VideoComment
	err := DB.Where("author_id = ?", userId).Order("comment_id desc").
		First(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}
