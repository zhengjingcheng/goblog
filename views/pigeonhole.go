package views

import (
	"net/http"
	"zjc_blog/common"
	"zjc_blog/service"
)

func (*HTMLApi) Pigeonhole(w http.ResponseWriter,r *http.Request)  {
	pigeonhole := common.Template.Pigeonhole

	pigeonholeRes := service.FindPostPigeonhole()
	pigeonhole.WriteData(w,pigeonholeRes)
}
