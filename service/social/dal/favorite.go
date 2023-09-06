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

func GetFavoriteVideoListByUid(ctx context.Context,
	userID int64) ([]int64, error) {
	var favoriteIDList []int64
	// 获取用户的点赞列表
	if err := DB.WithContext(ctx).Model(&model.UserFavoriteVideo{}).Where(
		"user_id = ?", userID).Pluck("favorite_id", &favoriteIDList).Error; err != nil {
		log.Errorw("get favorite list failed", "err", err)
		return nil, err
	}

	return favoriteIDList, nil
}
func GetFavoriteVideoSetByUid(ctx context.Context,
	userID int64) (map[int64]struct{}, error) {
	var favoriteIDList []int64
	favoriteIDSet := make(map[int64]struct{})
	// 获取用户的点赞列表
	if err := DB.WithContext(ctx).Model(&model.UserFavoriteVideo{}).Where(
		"user_id = ?", userID).Pluck("favorite_id", &favoriteIDList).Error; err != nil {
		log.Errorw("get favorite list failed", "err", err)
		return nil, err
	}
	// 将点赞列表转换为集合
	for _, v := range favoriteIDList {
		favoriteIDSet[v] = struct{}{} // 新增
	}

	return favoriteIDSet, nil
}

func GetFavoriteVideoIDListByUid(ctx context.Context,
	userID int64) ([]int64, error) {
	var favoriteIDList []int64
	// 获取用户的点赞列表
	if err := DB.WithContext(ctx).Model(&model.UserFavoriteVideo{}).Where(
		"user_id = ?", userID).Pluck("favorite_id", &favoriteIDList).Error; err != nil {
		log.Errorw("get favorite list failed", "err", err)
		return nil, err
	}

	return favoriteIDList, nil
}
