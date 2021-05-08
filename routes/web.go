package routes

import (
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
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show") // 文章详情
	r.HandleFunc("/articles", ac.Index).Methods("GET").Name("articles.index")           // 文章首页

	r.HandleFunc("/articles/create", ac.Create).Methods("GET").Name("articles.create") // 文章创建
	r.HandleFunc("/articles", ac.Store).Methods("POST").Name("articles.store")         // 文章保存

	r.HandleFunc("/articles/{id:[0-9]+}/edit", ac.Edit).Methods("GET").Name("articaled.edit") // 文章编辑页面

	r.HandleFunc("/articles/{id:[0-9]+}", ac.Update).Methods("POST").Name("articles.update") // 文章编辑

	r.HandleFunc("/articles/{id:[0-9]+}/delete", ac.Delete).Methods("POST").Name("articles.delete") //文章删除

	// 用户认证
	auc := new(controllers.AuthController)
	r.HandleFunc("/auth/register", auc.Register).Methods("GET").Name("auth.register")
	r.HandleFunc("/auth/do-register", auc.DoRegister).Methods("POST").Name("auth.doregister")

	// 静态资源
	r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))

	// 中间件：强制内容类型为 HTML
	// r.Use(middlewares.ForceHTML)
}
