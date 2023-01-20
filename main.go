package main

import (
	"ByteTechTraining/settings"
	"fmt"
)

func main() {
	// 初始化数据库
	err := settings.InitDatabase()
	if err != nil {
		panic(err)
		return
	}

	// 初始化gin引擎
	server, err := settings.InitGinEngine()
	if err != nil {
		return
	}

	// 项目启动
	err = server.Run(":8080")
	if err != nil {
		fmt.Println("Engine start error")
		return
	}

}
