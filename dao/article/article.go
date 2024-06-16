package article

import (
	"errors"
	"github.com/Tianyi822/easy_blog_server/db"
	"github.com/Tianyi822/easy_blog_server/logger"
	"time"
)

// ArticleDao 文章数据访问对象
type ArticleDao struct {
	Aid        string    `gorm:"primaryKey;size:36"`
	Title      string    `gorm:"size:50"`
	BriefIntro string    `gorm:"size:100"`
	Url        string    `gorm:"size:100"`
	Salt       string    `gorm:"size:32"`
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
}

// TableName 设置表名
func (ad *ArticleDao) TableName() string {
	return "EASY_BLOG_ARTICLE"
}

// Create 创建文章
func (ad *ArticleDao) Create() (err error) {
	// 开启事务
	tx := db.SqliteDb.Begin()
	defer func() {
		if r := recover(); r != nil {
			logger.Error("创建文章记录失败: %v", r)
			tx.Rollback()
		}
	}()

	// 插入文章记录
	logger.Info("插入文章记录...")
	result := tx.Create(ad)

	if result.Error != nil {
		tx.Rollback()
		msg := "插入文章记录失败: " + result.Error.Error()
		logger.Error(msg)
		return errors.New(msg)
	} else {
		tx.Commit()
		return nil
	}
}
