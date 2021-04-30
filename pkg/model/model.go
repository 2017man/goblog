package model

import (
	"goblog/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

// ConnectDB 初始化模型
func ConnectDB() *gorm.DB {
	var err error
	config := mysql.New(mysql.Config{
		DSN: "root:root@tcp(127.0.0.1:3306)/goblog?charset=utf8mb4&parseTime=True&loc=Local",
	})

	//准备数据库连接池
	DB, err = gorm.Open(config, &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	})

	logger.LogError(err)

	return DB
}
