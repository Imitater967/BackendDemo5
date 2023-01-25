package models

import "time"

type VideoModel struct {
	Id       int64
	Uploader int64
	Title    string
	Date     time.Time
	Deleted  int
}

func (*VideoModel) TableName() string {
	return "video"
}
