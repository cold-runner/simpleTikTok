package dal

import (
	"context"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/relation/dal/model"
	"github.com/cold-runner/simpleTikTok/service/relation/dal/redis"
	"gorm.io/gorm"
)

// CreateFollowRelation 创建用户关注关系
func CreateFollowRelation(ctx context.Context, userID, followingID int64) error {
	relation := &model.UserFollowRelation{UserID: userID,
		FollowingID: followingID}
	err := DB.WithContext(ctx).Create(relation).Error
	if err != nil {
		log.Errorw("create user relation failed", "err", err)
		return err
	}

	return nil
}

// GetFollowListByUserID  获取用户关注的用户列表
func GetFollowListByUserID(ctx context.Context, userID int64) ([]int64, error) {
	log.Debugw("get follow list by user id in dal", "userID", userID)
	var relations []model.UserFollowRelation
	err := DB.WithContext(ctx).Where("user_id = ?", userID).Find(&relations).Error
	if err != nil {
		log.Errorw("get follow list failed", "err", err)
		return nil, err
	}

	followingIDs := make([]int64, len(relations))
	for i, relation := range relations {
		followingIDs[i] = relation.FollowingID
	}
	log.Debugw("get follow list by user id in dal success", "followingIDs", followingIDs)
	return followingIDs, nil
}

// GetFollowerListByUserID 获取用户的粉丝列表
func GetFollowerListByUserID(ctx context.Context, userID int64) ([]int64, error) {
	var relations []model.UserFollowRelation
	err := DB.WithContext(ctx).Where("following_id = ?", userID).Find(&relations).Error
	if err != nil {
		log.Errorw("get follower list failed", "err", err)
		return nil, err
	}

	followerIDs := make([]int64, len(relations))
	for i, relation := range relations {
		followerIDs[i] = relation.UserID
	}

	return followerIDs, nil
}

// QueryFollowing 检查用户是否关注了另一个用户
func QueryFollowing(ctx context.Context, userID, followingID int64) (bool, error) {
	var relation model.UserFollowRelation

	// 查询数据库以检查关系是否存在
	err := DB.WithContext(ctx).Where("user_id = ? AND following_id = ?", userID, followingID).First(&relation).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil // 没有找到关系，返回false
		}
		log.Errorw("check user following failed", "err", err)
		return false, err // 查询时出现错误，返回错误
	}

	return true, nil // 找到关系，返回true
}

// DeleteFollowRelation 删除用户关注关系
func DeleteFollowRelation(ctx context.Context, userID, followingID int64) error {
	// 定义一个关系对象，用于执行删除
	relation := model.UserFollowRelation{
		UserID:      userID,
		FollowingID: followingID,
	}

	// 执行redis缓存删除操作
	redis.RemoveRelationCache(userID, followingID)

	err := DB.WithContext(ctx).Delete(&relation).Error
	if err != nil {
		log.Errorw("delete user relation failed", "err", err)
		return err // 如果删除时出现错误，则返回错误
	}

	return nil // 删除成功
}

func GetUserFollowSetAndList(ctx context.Context,
	uid int64) (map[int64]struct{}, []int64, error) {
	followSet := make(map[int64]struct{})
	followList, err := GetFollowListByUserID(ctx, uid)
	if err != nil {
		log.Errorw("get follow set failed", "err", err)
		return nil, nil, err
	}
	// 将关注列表转换为集合
	for _, v := range followList {
		followSet[v] = struct{}{}
	}
	return followSet, followList, nil
}
