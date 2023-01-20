package daos

import (
	"ByteTechTraining/globals/database"
	"ByteTechTraining/models"
	"errors"
)

type UserAuthDao struct {
	models.UserAuthModel
}

func (m *UserAuthDao) Get() error {
	mysqlManage := database.GetMysqlClient()
	return mysqlManage.Where("id", m.Id).Find(m).Error
}
func (m *UserAuthDao) Add() error {
	mysqlManage := database.GetMysqlClient()
	err := m.Get()
	if err == nil {
		return errors.New("数据已存在")
	}
	// 创建一条评论记录并返回error信息
	return mysqlManage.Create(&m).Error
}
