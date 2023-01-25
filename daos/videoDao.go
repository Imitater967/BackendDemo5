package daos

import (
	"ByteTechTraining/globals/database"
	"ByteTechTraining/models"
	"gorm.io/gorm"
	"time"
)

type VideoDao struct {
	models.VideoModel
	tx *gorm.DB `gorm:"-"`
}

// 上传

func (m *VideoDao) PreUpload() error {
	mysqlManage := database.GetMysqlClient()
	m.tx = mysqlManage.Begin()
	m.Date = time.Now() // 创建一条评论记录并返回error信息
	return m.tx.Create(&m).Error
}
func (m *VideoDao) CancelUpload() {
	m.tx.Rollback()
}
func (m *VideoDao) FinishUpload() error {
	return m.tx.Commit().Error
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
