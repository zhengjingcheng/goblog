package views

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"zjc_blog/common"
	"zjc_blog/service"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request)  {
	detail := common.Template.Detail
	//获取路径参数
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path,"/p/")
	//7.html
	pIdStr = strings.TrimSuffix(pIdStr,".html")
	pid,err := strconv.Atoi(pIdStr)//获取文章id
	if err != nil {
		detail.WriteError(w,errors.New("不识别此请求路径"))
		return
	}
	postRes,err := service.GetPostDetail(pid)//返回文章全文
	if err != nil {
		detail.WriteError(w,errors.New("查询出错"))
		return
	}
	detail.WriteData(w,postRes)

}
