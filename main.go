package main

import (
	"fmt"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello, 这里是 goblog</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>"+
			"<p>如有疑惑，请联系我们。</p>")
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
}

func main() {
	router := http.NewServeMux()
	// 主页
	router.HandleFunc("/", defaultHandler)
	// 关于页
	// 列表 or 创建
    router.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            fmt.Fprint(w, "访问文章列表")
        case "POST":
            fmt.Fprint(w, "创建新的文章")
        }
    })
	// 关于页
	router.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":4000", router)
}
