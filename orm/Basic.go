package orm

import "gorm.io/gorm"

type UserData struct {
	Username string
	Password string
	UserId   uint `gorm:"primaryKey"`
	Token    string
	Expire   uint
}
type UserInfo struct {
	gorm.Model
	UserId   uint `gorm:"primaryKey"`
	Nickname string
	Follow   []*UserInfo `gorm:"foreignKey:UserId"`
	Follower []*UserInfo `gorm:"foreignKey:UserId"`
}

func (v UserInfo) TableName() string {
	return "user_info"
}

type Video struct {
	VideoId uint
}
