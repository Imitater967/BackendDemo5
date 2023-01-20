package models

import "time"

type UserAuthModel struct {
	Name     string
	Password string
	Id       string
	Token    string
	Expire   time.Time
}

func (*UserAuthModel) TableName() string {
	return "user_auth"
}
