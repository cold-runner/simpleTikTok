package dal

import (
	"context"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/social/dal/model"
)

func AddFavorite(ctx context.Context, userID int64, favoriteID int64) error {
	favorite := model.UserFavoriteVideo{
		UserID:     userID,
		FavoriteID: favoriteID,
	}
	// 创建一条新的点赞记录
	if err := DB.WithContext(ctx).Create(&favorite).Error; err != nil {
		log.Errorw("add favorite failed", "err", err)
		return err
	}
	return nil
}

// RemoveFavorite 用于删除一条点赞数据
func RemoveFavorite(ctx context.Context, userID int64,
	favoriteID int64) error {
	favorite := model.UserFavoriteVideo{
		UserID:     userID,
		FavoriteID: favoriteID,
	}
	// 删除点赞记录
	if err := DB.WithContext(ctx).Delete(&favorite).Error; err != nil {
		log.Errorw("remove favorite failed", "err", err)
		return err
	}
	return nil
}
