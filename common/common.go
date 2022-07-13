package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"zjc_blog/config"
	"zjc_blog/models"
)

var Template models.HtmlTemplate

func LoadTemplate()  {
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		//耗时
		var err error
		Template,err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(any(err))
		}
		w.Done()
	}()
	w.Wait()
}

func Success(w http.ResponseWriter,data interface{}){
	var result models.Result
	result.Code = 200
	result.Error = ""
	result.Data = data
	resultJson,_ := json.Marshal(result)
	w.Header().Set("Content-Type","application/json")
	_,err := w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}
func Error(w http.ResponseWriter,err error){
	var result models.Result
	result.Code = -999
	result.Error = err.Error()
	resultJson,_ := json.Marshal(result)
	w.Header().Set("Content-Type","application/json")
	_, err = w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}

func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &params)
	return params
}