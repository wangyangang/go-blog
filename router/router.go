package router

import (
	"go-blog/api"
	"go-blog/views"
	"net/http"
)

func Router() {

	http.HandleFunc("/index.html", views.HTML.Index)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	// 处理静态资源
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/ "))))
}
