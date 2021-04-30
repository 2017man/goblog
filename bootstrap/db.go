package bootstrap

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"time"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() {
	// 建立数据库连接池
	db := model.ConnectDB()

	// 命令行打印数据库请求的信息
	sqlDb, _ := db.DB()

	// 设置最大连接数
	sqlDb.SetMaxOpenConns(25)
	// 设置最大空闲连接数
	sqlDb.SetMaxIdleConns(25)
	// 设置每个链接的过期时间
	sqlDb.SetConnMaxLifetime(5 * time.Minute)

	// 尝试连接，失败会报错
	error := sqlDb.Ping()
	logger.LogError(error)

}
