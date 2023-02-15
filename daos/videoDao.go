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
}
func GetUploadedVideos(uploader int64) ([]*VideoDao, error) {
	var videos []*VideoDao
	mysqlManager := database.GetMysqlClient()
	db := mysqlManager.Model(&VideoDao{}).Where("uploader", uploader).Find(&videos)
	return videos, db.Error
}
func QueryFeed(date time.Time) ([]*VideoDao, time.Time, error) {
	var videos []*VideoDao
	mysqlManager := database.GetMysqlClient()
	//select * from test where create_time >= '投稿时间' order by cast(date as datetime) desc
	db := mysqlManager.Model(&VideoDao{}).Limit(30).Order("date desc").Where("date > ?", date).Find(&videos)
	var lastVideo VideoDao
	db.Last(&lastVideo)
	return videos, lastVideo.Date, db.Error
}

// Delete 删
func (m *VideoDao) Delete() error {
	mysqlManage := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManage.Model(&m).Update("deleted", 1).Error
}
