package models

import "time"

type CommentModel struct {
	Id      int64
	UserId  int64
	VideoId int64
	Content string
	Deleted int
	Date    time.Time
}

func (*CommentModel) TableName() string {
	return "comment"
}
