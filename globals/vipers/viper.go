// Package vipers Viper 是一个完整的 Go 应用程序配置解决方案
package vipers

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"sync"
)

var (
	databaseViper     *viper.Viper
	databaseOnceViper sync.Once
)

// GetDatabaseViper 使用单例模式读取configs/database.yaml
func GetDatabaseViper() *viper.Viper {
	databaseOnceViper.Do(func() {
		databaseViper = viper.New()
		databaseViper.SetConfigName("database")  // 配置文件名称，无扩展名
		databaseViper.AddConfigPath("./configs") // 添加搜索路径
		databaseViper.SetConfigType("yaml")

		err := databaseViper.ReadInConfig() // 查找并读取配置文件
		if err != nil {
			panic(err)
		}

		databaseViper.WatchConfig() // 监听配置文件的修改并重新读取
		databaseViper.OnConfigChange(func(e fsnotify.Event) {
			// viper发生更改时调用该函数
		})
	})

	return databaseViper
}
