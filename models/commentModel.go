package models

import "time"

type CommentModel struct {
	CommentID   int64     `gorm:"column:comment_id;primaryKey;type:bigint;not null;"`
	CommenterID int64     `gorm:"column:commenter_id;type:bigint(22);"`
	Content     string    `gorm:"column:content;type:varchar(2000);"`
	IsDeleted   int       `gorm:"column:is_deleted;type:tinyint(1);default:0"`
	CreateTime  time.Time `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP"`
}

func (CommentModel) TableName() string {
	return "comment"
}
