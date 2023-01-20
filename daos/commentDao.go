package daos

import (
	"ByteTechTraining/globals/database"
	"ByteTechTraining/models"
)

type CommentDao struct {
	models.CommentModel
}

// Get 查
func (m *CommentDao) Get() error {
	mysqlManage := database.GetMysqlClient()
	return mysqlManage.Where("id", m.Id).Where("deleted", 0).Find(m).Error
}

// Add 增
func (m *CommentDao) Add() error {
	mysqlManage := database.GetMysqlClient()
	// 创建一条评论记录并返回error信息
	return mysqlManage.Create(&m).Error
}

// Update 改
func (m *CommentDao) Update(args map[string]any) error {
	mysqlManage := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManage.Model(&m).Updates(args).Error
}

// Delete 删
func (m *CommentDao) Delete() error {
	mysqlManage := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManage.Model(&m).Update("deleted", 1).Error
}
