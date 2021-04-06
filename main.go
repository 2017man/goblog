package main

import (
	"fmt",
	"net/http"
)

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, 这里是 goblog</h1>")
}

func main()  {
	http.HandleFunc("/",HandlerFunc)
	http.ListenAndServer(":3000",nil)
}
