package settings

import (
	"ByteTechTraining/globals/database"
	"fmt"
)

func InitDatabase() (err error) {

	err = database.InitMysqlClient()
	if err != nil {
		fmt.Println("mysql init error: ", err)
		return
	}

	return nil
}
