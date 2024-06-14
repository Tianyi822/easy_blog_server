package main

import (
	"flag"
	"github.com/Tianyi822/easy_blog_server/config"
	"github.com/Tianyi822/easy_blog_server/db"
	"github.com/Tianyi822/easy_blog_server/logger"
)

// parseArgs 解析命令行参数
func parseArgs() string {
	// 配置文件路径
	var configFilePath string

	// 解析命令行参数
	flag.StringVar(&configFilePath, "conf-path", "resources/config/config-dev.yaml", "config file path")

	flag.Parse()

	return configFilePath
}

// initServer 初始化服务
func initServer(configPath string) {
	// 加载配置文件
	config.LoadConfig(configPath)
	// 初始化日志组件
	logger.InitLogger()
	// 连接 Sqlite 数据库
	db.ConnectSqlite()
}

func main() {
	initServer(parseArgs())
	println(config.ServerConf.Port)
}
