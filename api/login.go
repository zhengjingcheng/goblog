package api

import (
	"net/http"
	"zjc_blog/common"
	"zjc_blog/service"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request)  {
	//接收用户名和密码 返回 对应的json数据
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	loginRes,err := service.Login(userName,passwd)
	if err != nil {
		common.Error(w,err)
		return
	}
	common.Success(w,loginRes)
}
