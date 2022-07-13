package router

import (
	"net/http"
	"zjc_blog/api"
	"zjc_blog/views"
)

func Router() {
	//1.页面 views 2.数据 api（json）3.静态资源
	http.HandleFunc("/", views.HTML.Index)//主界面
	http.HandleFunc("/c/",views.HTML.Category)//分类界面
	http.HandleFunc("/login",views.HTML.Login)//登录界面
	http.HandleFunc("/api/v1/login",api.API.Login)//登录接口
	http.HandleFunc("/p/",views.HTML.Detail)//文章详情
	http.HandleFunc("/writing",views.HTML.Writing)//写作界面
	http.HandleFunc("/api/v1/post",api.API.SaveAndUpdatePost)//更新和保存
	http.HandleFunc("/api/v1/post/",api.API.GetPost)//编辑文章
	http.HandleFunc("/pigeonhole",views.HTML.Pigeonhole)//归档页面
	http.HandleFunc("/api/v1/post/search",api.API.SearchPost)//搜索界面
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}