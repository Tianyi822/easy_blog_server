package db

import (
	"github.com/Tianyi822/easy_blog_server/config"
	"github.com/Tianyi822/easy_blog_server/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"sync"
)

var SqliteDb *gorm.DB

var sqliteOnceLock sync.Once

// ConnectSqlite 连接 Sqlite 数据库
func ConnectSqlite() {
	sqliteOnceLock.Do(
		func() {
			logger.Info("准备连接 Sqlite 数据库...")

			// 检查 sqlite 的数据路径是否存在
			if _, err := os.Stat(config.SqliteConf.DataPath); os.IsNotExist(err) {
				msg := "sqlite 数据路径不存在: " + config.SqliteConf.DataPath
				logger.Warn(msg)
				// 创建 sqlite 数据路径
				err := os.MkdirAll(config.SqliteConf.DataPath, os.ModePerm)
				if err != nil {
					msg := "创建 sqlite 数据路径失败: " + err.Error()
					logger.Error(msg)
					panic(msg)
				}
			}

			// 连接数据库
			db, err := gorm.Open(
				sqlite.Open(config.SqliteConf.DataPath+"/"+config.SqliteConf.DB+".db"),
				&gorm.Config{},
			)
			if err != nil {
				msg := "连接 sqlite 发生错误: " + err.Error()
				logger.Error(msg)
				panic(msg)
			}

			// 设置连接池
			sqlDB, err := db.DB()
			if err != nil {
				msg := "设置 sqlite 连接池错误: " + err.Error()
				logger.Error(msg)
				panic(msg)
			}

			// 设置最大连接数
			sqlDB.SetMaxOpenConns(config.SqliteConf.MaxOpen)
			// 设置最大空闲连接数
			sqlDB.SetMaxIdleConns(config.SqliteConf.MaxIdle)

			SqliteDb = db
		},
	)
}
