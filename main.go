package main

import (
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/pkg/database"
	"goblog/pkg/route"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var router = route.Router

func main() {
	//数据库初始化
	database.Initiate()
	//初始化
	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
