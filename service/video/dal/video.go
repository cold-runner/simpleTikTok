package dal

import (
	"context"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/video/dal/model"
	"gorm.io/gorm"
	"time"
)

func CreateVideo(ctx context.Context, video *model.Video) error {
	err := DB.WithContext(ctx).Create(video).Error
	if err != nil {
		log.Errorw("create video failed", "err", err)
		return err
	}
	return nil
}

// IncrementFavoriteCount 增加FavoriteCount
func IncrementFavoriteCount(ctx context.Context, videoID int64) error {
	err := DB.WithContext(ctx).Model(&model.Video{}).Where("video_id = ?", videoID).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
	if err != nil {
		log.Errorw("increment FavoriteCount failed", "err", err)
		return err
	}
	return nil
}

// DecrementFavoriteCount 减少FavoriteCount
func DecrementFavoriteCount(ctx context.Context, videoID int64) error {
	err := DB.WithContext(ctx).Model(&model.Video{}).Where("video_id = ?", videoID).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
	if err != nil {
		log.Errorw("decrement FavoriteCount failed", "err", err)
		return err
	}
	return nil
}

// IncrementCommentCount 增加CommentCount
func IncrementCommentCount(ctx context.Context, videoID int64) error {
	err := DB.WithContext(ctx).Model(&model.Video{}).Where("video_id = ?", videoID).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error
	if err != nil {
		log.Errorw("increment CommentCount failed", "err", err)
		return err
	}
	return nil
}

// DecrementCommentCount 减少CommentCount
func DecrementCommentCount(ctx context.Context, videoID int64) error {
	err := DB.WithContext(ctx).Model(&model.Video{}).Where("video_id = ?", videoID).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).Error
	if err != nil {
		log.Errorw("decrement CommentCount failed", "err", err)
		return err
	}
	return nil
}

func GetVideoByID(ctx context.Context, id int64) (*model.Video, error) {
	var video model.Video
	if err := DB.WithContext(ctx).Where("video_id = ?",
		id).First(&video).Error; err != nil {
		log.Errorw("get video by id failed", "err", err)
		return nil, err
	}
	return &video, nil
}

func GetVideosBeforeTimeWithAuthors(ctx context.Context,
	beforeTimeStamp int64) ([]*model.Video, []int64, error) {
	var videos []*model.Video
	var authorIDs []int64
	// 使用GORM进行查询，选取早于传入时间的视频，按创建时间降序排列，最多返回30个结果。
	err := DB.WithContext(ctx).Where("updated_at <= ?",
		time.UnixMilli(beforeTimeStamp)).Order("updated_at desc").
		Limit(30).Find(&videos).Error
	if err != nil {
		log.Errorw("mysql query videos failed", "err", err)
		return nil, nil, err
	}
	// 将作者的 ID 也放入结果中返回
	for _, video := range videos {
		authorIDs = append(authorIDs, video.AuthorID)
	}

	return videos, authorIDs, nil
}

// GetVideoListByUserId 通过用户ID获取视频列表
func GetVideoListByUserId(ctx context.Context,
	userId int64) ([]*model.Video, error) {
	var videos []*model.Video
	err := DB.WithContext(ctx).Where("author_id = ?", userId).Find(&videos).Error
	if err != nil {
		log.Errorw("mysql query user videos failed", "err", err)
		return nil, err
	}
	return videos, nil
}

// MGetVideoListByIds 通过视频ID列表获取视频列表
func MGetVideoListByIds(ctx context.Context, videoIds []int64) ([]*model.
	Video, error) {
	var videos []*model.Video
	err := DB.WithContext(ctx).Where("video_id in ?", videoIds).Find(&videos).Error
	if err != nil {
		log.Errorw("mysql query videos by ids failed", "err", err)
		return nil, err
	}
	return videos, nil
}

func GetAuthorIdByVideoId(ctx context.Context, videoId int64) (int64, error) {
	var video *model.Video
	// 在数据库中查找对应的记录
	if err := DB.WithContext(ctx).First(&video, videoId).Error; err != nil {
		// 如果出现错误（例如，找不到记录），则返回错误
		return 0, err
	}

	// 返回找到的 AuthorID
	return video.AuthorID, nil
}
