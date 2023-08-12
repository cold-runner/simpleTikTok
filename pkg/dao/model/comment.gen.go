// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameComment = "comment"

// Comment mapped from table <comment>
type Comment struct {
	ID        int64     `gorm:"column:id;type:int;primaryKey" json:"id,string"`
	UserID    int64     `gorm:"column:user_id;type:int" json:"user_id"`
	Content   string    `gorm:"column:content;type:text" json:"content"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
	VedioID   int64     `gorm:"column:vedio_id;type:int" json:"vedio_id"`
}

// TableName Comment's table name
func (*Comment) TableName() string {
	return TableNameComment
}
