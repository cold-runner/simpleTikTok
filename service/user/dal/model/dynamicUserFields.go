package model

// DynamicUserFields 用于保存用户信息当中的动态字段
type DynamicUserFields struct {
	FollowCount    int64 `json:"follow_count"`
	FollowerCount  int64 `json:"follower_count"`
	TotalFavorited int64 `json:"total_favorited"`
	WorkCount      int64 `json:"work_count"`
	FavoriteCount  int64 `json:"favorite_count"`
}
