package orm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const mysqlAddress = "root:123456@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

var gormDB *gorm.DB

func ConnectToDatabase() {
	var err error
	gormDB, err = gorm.Open(mysql.Open(mysqlAddress), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
}
