package views

import (
	"net/http"
	"zjc_blog/common"
	"zjc_blog/config"
)

func (*HTMLApi) Login(w http.ResponseWriter,r *http.Request){
	login := common.Template.Login

	login.WriteData(w,config.Cfg.Viewer)
}