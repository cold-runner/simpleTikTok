package dal

import (
	"context"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/user/dal/model"
	"gorm.io/gorm"
)

// CreateUser create user info
func CreateUser(ctx context.Context, user *model.User) error {
	err := DB.WithContext(ctx).Create(user).Error
	if err != nil {
		log.Errorw("create user failed", "err", err)
		return err
	}
	return nil
}

// GetUserByUserName 1 - 1 通过username获取 User 对象
func GetUserByUserName(ctx context.Context,
	userName string) (*model.User, error) {
	res := &model.User{}
	if err := DB.WithContext(ctx).Where("username = ?",
		userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetUserByUserId 1 - 1 通过 id 获取 User 对象
func GetUserByUserId(ctx context.Context, uid int64) (*model.User, error) {
	res := &model.User{}
	// 指定查询字段，排除password
	if err := DB.WithContext(ctx).Select("id, username, follow_count, follower_count, avatar, background_image, signature, total_favorited, work_count, favorite_count").Where("id = ?", uid).Find(&res).Error; err != nil {
		log.Errorw("get user by user id failed", "err", err)
		return nil, err
	}
	log.Debugw("get user static fields success", "user", res)
	return res, nil
}

func GetDynamicUserFields(ctx context.Context,
	uid int64) (*model.DynamicUserFields, error) {
	// 初始化用于保存查询结果的结构体
	user := &model.User{}

	// 执行数据库查询，选取需要的字段
	if err := DB.WithContext(ctx).Select("follow_count, follower_count, total_favorited, work_count, favorite_count").Where("id = ?", uid).Find(user).Error; err != nil {
		log.Errorw("get dynamic user fields failed", "err", err)
		return nil, err
	}
	log.Debugw("get dynamic user fields success", "user", user)
	return &model.DynamicUserFields{
		FollowCount:    user.FollowCount,
		FollowerCount:  user.FollowerCount,
		TotalFavorited: user.TotalFavorited,
		WorkCount:      user.WorkCount,
		FavoriteCount:  user.FavoriteCount,
	}, nil
}

// IncreaseUserFollowCount 用于增加用户的关注数
func IncreaseUserFollowCount(ctx context.Context, userID int64) error {
	if err := DB.WithContext(ctx).Model(&model.User{}).Where("id = ?", userID).
		UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
		log.Errorw("increase user follow count failed", "err", err)
		return err
	}
	return nil
}

// DecreaseUserFollowCount 用于减少用户的关注数
func DecreaseUserFollowCount(ctx context.Context, userID int64) error {
	if err := DB.WithContext(ctx).Model(&model.User{}).Where("id = ?", userID).
		UpdateColumn("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
		log.Errorw("decrease user follow count failed", "err", err)
		return err
	}
	return nil
}

// IncreaseUserFollowerCount 用于增加用户的粉丝数
func IncreaseUserFollowerCount(ctx context.Context, toUid int64) error {
	if err := DB.WithContext(ctx).Model(&model.User{}).Where("id = ?", toUid).
		UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
		log.Errorw("increase user follower count failed", "err", err)
		return err
	}
	return nil
}

// DecreaseUserFollowerCount 用于减少用户的粉丝数
func DecreaseUserFollowerCount(ctx context.Context, toUid int64) error {
	if err := DB.WithContext(ctx).Model(&model.User{}).Where("id = ?", toUid).
		UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
		log.Errorw("decrease user follower count failed", "err", err)
		return err
	}
	return nil
}

// IncreaseUserTotalFavorited 用于增加用户的被赞数
func IncreaseUserTotalFavorited(ctx context.Context, userID int64) error {
	if err := DB.WithContext(ctx).Model(&model.User{}).Where("id = ?", userID).
		UpdateColumn("total_favorited", gorm.Expr("total_favorited + ?", 1)).Error; err != nil {
		log.Errorw("increase user total favorited count failed", "err", err)
		return err
	}
	return nil
}

// DecreaseUserTotalFavorited 用于减少用户的被赞数
func DecreaseUserTotalFavorited(ctx context.Context, userID int64) error {
	if err := DB.WithContext(ctx).Model(&model.User{}).Where("id = ?", userID).
		UpdateColumn("total_favorited", gorm.Expr("total_favorited - ?", 1)).Error; err != nil {
		log.Errorw("decrease user total favorited count failed", "err", err)
		return err
	}
	return nil
}

// IncreaseUserWorkCount 用于增加用户的作品数
func IncreaseUserWorkCount(ctx context.Context, userID int64) error {
	if err := DB.WithContext(ctx).Model(&model.User{}).Where("id = ?", userID).
		UpdateColumn("work_count", gorm.Expr("work_count + ?", 1)).Error; err != nil {
		log.Errorw("increase user work count failed", "err", err)
		return err
	}
	return nil
}

// DecreaseUserWorkCount 用于减少用户的作品数
func DecreaseUserWorkCount(ctx context.Context, userID int64) error {
	if err := DB.WithContext(ctx).Model(&model.User{}).Where("id = ?", userID).
		UpdateColumn("work_count", gorm.Expr("work_count - ?", 1)).Error; err != nil {
		log.Errorw("decrease user work count failed", "err", err)
		return err
	}
	return nil
}

// IncreaseUserFavoriteCount 用于增加用户的点赞数
func IncreaseUserFavoriteCount(ctx context.Context, userID int64) error {
	if err := DB.WithContext(ctx).Model(&model.User{}).Where("id = ?", userID).
		UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
		log.Errorw("increase user favorite count failed", "err", err)
		return err
	}
	return nil
}

// DecreaseUserFavoriteCount 用于减少用户的点赞数
func DecreaseUserFavoriteCount(ctx context.Context, userID int64) error {
	if err := DB.WithContext(ctx).Model(&model.User{}).Where("id = ?", userID).
		UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
		log.Errorw("decrease user favorite count failed", "err", err)
		return err
	}
	return nil
}
