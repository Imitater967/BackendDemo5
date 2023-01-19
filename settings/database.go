package settings

import (
	"ByteTechTraining/globals/database"
	"fmt"
	"github.com/spf13/viper"
)

func InitDatabase() (err error) {
	if viper.GetBool("system.UseMysql") {
		err = database.InitMysqlClient()
		if err != nil {
			fmt.Println("mysql init error: ", err)
			return
		}
	}
	return
}
