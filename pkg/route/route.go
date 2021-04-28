package route

import (
	"goblog/routes"

	"github.com/gorilla/mux"
)

// Router 路由对象
var Router *mux.Router

//初始化路由
func Initiate() {
	Router = mux.NewRouter()
	routes.RegisterWebRoutes(Router)
}

// RouteName2URL 通过路由名称来获取 URL
func Name2URL(routeName string, pairs ...string) string {
	url, err := Router.Get(routeName).URL(pairs...)
	if err != nil {
		// checkError(err)
		return ""
	}

	return url.String()
}
