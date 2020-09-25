package app

import (
	"flag"
	"fmt"
	"kitstart_goweb/app/utils"
)

// 参数值
var configDir string

// InitApp - 初始化App 整个项目
// 加载配置
// 初始化路由表
func InitApp() {
	// 参数配置
	flag.StringVar(&configDir, "c", "", "config directory")

	flag.Parse()

	if configDir == "" {
		fmt.Println("config directory is required to start app")
		return
	}

	err := utils.LoadConfig(configDir)

	if err != nil {
		fmt.Println(err)
	}

	port, err := utils.ConfigInt("app.port")
	host, err := utils.ConfigString("database.dev.host")
	fmt.Println("port is ", port, host, err)

}
