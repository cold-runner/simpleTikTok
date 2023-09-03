package model

const TableNameRelation = "user_follow_relation"

// UserFollowRelation mapped from table <user_follow_relation>
type UserFollowRelation struct {
	UserID      int64 `gorm:"primaryKey;column:user_id"`
	FollowingID int64 `gorm:"primaryKey;column:following_id"`
}

// TableName User's table name
func (*UserFollowRelation) TableName() string {
	return TableNameRelation
}
