package daos

import (
	"ByteTechTraining/globals/database"
	"ByteTechTraining/models"
	"time"
)

type VideoDao struct {
	models.VideoModel
}

// 上传
func (m *VideoDao) Upload() error {
	mysqlManage := database.GetMysqlClient()
	m.Date = time.Now() // 创建一条评论记录并返回error信息
	return mysqlManage.Create(&m).Error
}

// Get 查
func (m *VideoDao) Get() error {
	mysqlManage := database.GetMysqlClient()
	return mysqlManage.Where("id", m.Id).Where("deleted", 0).First(m).Error
} // Delete 删
func (m *VideoDao) Delete() error {
	mysqlManage := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManage.Model(&m).Update("deleted", 1).Error
}
