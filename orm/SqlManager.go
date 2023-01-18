package orm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

const mysqlAddress = "root:123456@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"

var db *gorm.DB

func ConnectToDatabase() {
	var err error
	db, err = gorm.Open(mysql.Open(mysqlAddress), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
}

// Basic
func UploadUserInfo() {
	user1 := UserInfo{UserId: 1, Nickname: "143"}
	var user2 = UserInfo{UserId: 1, Nickname: "144", Follow: []*UserInfo{&user1}, Follower: []*UserInfo{}}
	db.Create(user1)
	db.Create(user2)
}

// Interact
func UploadComment() {
	comment := Comment{VideoId: 1, CommentId: 1, Content: "LOL", CreateDate: time.Now()}
	db.Create(comment)
}
func GetComment() []*Comment {
	comments := make([]*Comment, 10)
	video := Video{VideoId: 1}
	db.Where("commit_id <> ?", video.VideoId).Limit(30).Find(&comments)
	var err = db.Error
	if err != nil {
		log.Println(err)
	}
	return comments
}
