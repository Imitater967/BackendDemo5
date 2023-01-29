package daos

import (
	"ByteTechTraining/globals/database"
	"ByteTechTraining/models"
	"errors"
)

type VideoFavoriteDao struct {
	models.VideoFavoriteModel
}

func (m *VideoFavoriteDao) Get() error {
	mysqlManage := database.GetMysqlClient()
	db := mysqlManage.Where("user", m.User).Where("video", m.Video).First(&m)
	return db.Error
}

// 标记为喜欢,也就是上传一条数据
func (m *VideoFavoriteDao) Mark() error {
	mysqlManager := database.GetMysqlClient()
	//查询记录无错误,也就是有记录
	favoriteExist := m.Get()
	if favoriteExist == nil {
		return errors.New("已有相关记录")
	}
	db := mysqlManager.Model(&VideoFavoriteDao{}).Create(&m)
	return db.Error
}

// Delete 删
func (m *VideoFavoriteDao) Unmark() error {
	mysqlManage := database.GetMysqlClient()
	db := mysqlManage.Where("user", m.User).Where("video", m.Video).Delete(&VideoFavoriteDao{})
	return db.Error
}

// 根据用户id,返回所有符合条件的喜欢记录
func GetFavoriteVideos(user int64) ([]*VideoFavoriteDao, error) {
	var videos []*VideoFavoriteDao
	mysqlManager := database.GetMysqlClient()
	db := mysqlManager.Where("user", user).Find(&videos)
	return videos, db.Error
}
