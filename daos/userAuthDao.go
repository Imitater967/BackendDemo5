package daos

import (
	"ByteTechTraining/globals/database"
	"ByteTechTraining/models"
	"ByteTechTraining/utils"
	"errors"
	"time"
)

type UserAuthDao struct {
	models.UserAuthModel
}

func (m *UserAuthDao) Get() error {
	if m.Token == "" {
		return errors.New("Token Can't Be Null")
	}
	mysqlManage := database.GetMysqlClient()
	return mysqlManage.Where("token", m.Token).First(&m).Error
}
func (m *UserAuthDao) Query() error {
	mysqlManage := database.GetMysqlClient()
	return mysqlManage.Where("name", m.Name).First(&m).Error
}
func (m *UserAuthDao) Register() error {
	mysqlManage := database.GetMysqlClient()
	err := m.Query()
	//没有则会返回错误,也就是如果有,就没有返回错误
	if err == nil {
		return errors.New("数据已存在")
	}
	m.Expire = time.Now().AddDate(0, 0, 7)
	// 创建一条评论记录并返回error信息
	return mysqlManage.Create(&m).Error
}

// 用户登录, 生成新的Token并刷新token时间
func (m *UserAuthDao) Login() error {
	sql := database.GetMysqlClient()
	//如果找不到会有报错,找得到则没有
	var tx = sql.Where("name", &m.Name).Where("password", &m.Password).First(&m)
	if tx.Error != nil {
		return errors.New("登录失败,用户名或密码错误")
	}
	token, err, exp := utils.GenerateToken((*m).UserAuthModel)
	if err != nil {
		return errors.New("token生成失败")
	}
	m.Token = token
	m.Expire = exp
	
	sql.Model(&m).Where("name", &m.Name).Updates(UserAuthDao{models.UserAuthModel{
		Token: m.Token, Expire: m.Expire}})
	return nil
}
