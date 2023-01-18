package orm

import (
	"time"
)

type Comment struct {
	//这个Video的外键是Comment.VideoId
	Video      Video `gorm:"foreignkey:VideoId"`
	VideoId    uint
	CommentId  uint `gorm:"primaryKey"`
	Content    string
	CreateDate time.Time
}

func (v Comment) TableName() string {
	return "comment"
}
