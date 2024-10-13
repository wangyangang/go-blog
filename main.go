package main

import (
	"go-blog/common"
	"go-blog/router"
	"log"
	"net/http"
)

func init() {
	// 模板加载
	common.LoadTemplate()
}
func main() {
	server := http.Server{
		Addr: ":8080",
	}
	// 路由
	router.Router()
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
