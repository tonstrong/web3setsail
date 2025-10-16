package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"myBlog/internal/app/myBlog/model"
)

func main() {
	initDB()
}

func initDB() {
	var err error
	// 连接SQLite数据库（实际项目可替换为MySQL等其他数据库）
	db, err := gorm.Open(sqlite.Open("identifier.sqlite"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 显示SQL日志
	})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 自动迁移表结构（创建不存在的表）
	err = db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	if err != nil {
		log.Fatalf("表结构迁移失败: %v", err)
	}
}
