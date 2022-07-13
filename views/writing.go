package views

import (
	"net/http"
	"zjc_blog/common"
	"zjc_blog/service"
)

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request)  {
	writing := common.Template.Writing
	wr := service.Writing()
	writing.WriteData(w,wr)
}
