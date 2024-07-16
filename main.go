package main

import (
	"flag"
	"fmt"
	"github.com/Tianyi822/easy_blog_server/config"
	"github.com/Tianyi822/easy_blog_server/db"
	"github.com/Tianyi822/easy_blog_server/logger"
	"github.com/Tianyi822/easy_blog_server/routers"
	article_routers "github.com/Tianyi822/easy_blog_server/routers/article-routers"
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

func runSever() {
	logger.Info("加载路由")
	routers.IncludeOpts(article_routers.Routers)
	logger.Info("加载路由完成")

	logger.Info("服务配置")
	r := routers.InitRouter()
	port := fmt.Sprintf(":%v", config.ServerConf.Port)
	logger.Info("服务配置完成")

	logger.Info("启动服务: %v", port)
	if err := r.Run(port); err != nil {
		logger.Panic("startup service failed, err:%v\n", err)
	}
	logger.Info("服务启动成功")
}

// initServer 初始化服务
func initServer() {
	// 解析命令行参数
	configPath := parseArgs()
	// 加载配置文件
	config.LoadConfig(configPath)
	// 初始化日志组件
	logger.InitLogger()
	// 连接 Sqlite 数据库
	db.ConnectSqlite()
}

func main() {
	// 初始化服务
	initServer()
	// 注册路由
	runSever()
}
