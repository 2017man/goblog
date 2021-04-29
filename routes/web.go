package routes

import (
	"fmt"
	"goblog/app/http/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(r *mux.Router) {

	//静态页面
	pc := new(controllers.PagesController)

	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")        // 主页
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about") // 关于页
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)             // 自定义 404 页面

	//文章相关
	ac := new(controllers.ArticlesController)
	fmt.Println("111")
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show") // 文章详情
}
