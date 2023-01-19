package database

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	mysqlClient *gorm.DB
)

// GetMysqlClient 获取mysql连接
func GetMysqlClient() *gorm.DB {
	if gin.Mode() == gin.ReleaseMode {
		return mysqlClient
	}
	return mysqlClient.Debug()
}

// InitMysqlClient 初始化mysql连接
func InitMysqlClient() (err error) {
	//v := vipers.GetDatabaseViper()
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//	v.GetString("mysql.username"),
	//	v.GetString("mysql.password"),
	//	v.GetString("mysql.host"),
	//	v.GetString("mysql.port"),
	//	v.GetString("mysql.database"),
	//)
	//mysqlClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	return
	//}
	return nil
}
