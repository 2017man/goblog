package controllers

import (
	"fmt"
	"net/http"
)

// PagesController 处理静态页面
type PagesController struct {
}

//Home 首页
func (*PagesController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, 欢迎来到 goblog！</h1>")
}

//About 关于页
func (*PagesController) About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>关于页</h1>")
}

//NotFound 页面
func (*PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>")
}
