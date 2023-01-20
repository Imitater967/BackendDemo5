package daos

import (
	"ByteTechTraining/globals/database"
	"ByteTechTraining/models"
	"errors"
	"time"
)

type UserAuthDao struct {
	models.UserAuthModel
}

func (m *UserAuthDao) Get() error {
	mysqlManage := database.GetMysqlClient()
	return mysqlManage.Where("name", m.Name).Find(m).Error
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

// 用户登录, 生成新的Token并刷新token时间
func (m *UserAuthDao) Login() error {
	sql := database.GetMysqlClient()
	var data = UserAuthDao{}
	//如果找不到会有报错,找得到则没有
	var tx = sql.Where("name", m.Name).Where("password", m.Password).Find(&data)
	if tx.Error != nil {
		errors.New("登录失败,用户名或密码错误")
	}
	m.Token = GenerateToken()
	m.Expire = time.Now()
	return nil
}

// 生成token,并从数据库中检测是否有相等的token
func GenerateToken() string {
	return "123456"
}