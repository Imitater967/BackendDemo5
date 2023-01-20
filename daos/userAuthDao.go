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
	return mysqlManage.Where("name", m.Name).First(&m).Error
}
func (m *UserAuthDao) Register() error {
	mysqlManage := database.GetMysqlClient()
	err := m.Get()
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
	var data = UserAuthDao{}
	//如果找不到会有报错,找得到则没有
	var tx = sql.Where("name", m.Name).Where("password", m.Password).First(&data)
	if tx.Error != nil {
		return errors.New("登录失败,用户名或密码错误")
	}
	m.GenerateToken()
	m.Expire = time.Now().AddDate(0, 0, 7)
	return nil
}

// 生成token,并从数据库中检测是否有相等的token
func (m *UserAuthDao) GenerateToken() {
	m.Token = "123456"
}
