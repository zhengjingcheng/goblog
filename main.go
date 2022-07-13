package main

import (
	"log"
	"net/http"
	"zjc_blog/common"
	"zjc_blog/router"
)

func init(){
	//加载模板
	common.LoadTemplate()
}
func main(){
	//程序入口 一个项目只有一个入口
	//web程序，http协议ip，port
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	//路由功能
	router.Router()
	if err:=server.ListenAndServe();err!=nil{
		log.Panicln(err)
	}
}
