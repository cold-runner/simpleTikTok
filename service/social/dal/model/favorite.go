package model

const TableNameFavorite = "user_favorite_video"

type UserFavoriteVideo struct {
	UserID     int64 `gorm:"primaryKey;column:user_id"`
	FavoriteID int64 `gorm:"primaryKey;column:favorite_id"`
}

func (*UserFavoriteVideo) TableName() string {
	return TableNameFavorite
}
