package reponse

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error, code int) {
	var body Body
	if err != nil {
		if code != 0 {
			body.Code = code
		} else {
			body.Code = -1
		}
		body.Msg = err.Error()
	} else {
		body.Msg = "OK"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}

//err != nil时输出data
func ResponseWithFailData(w http.ResponseWriter, resp interface{}, err error, code int) {
	var body Body
	if err != nil {
		if code != 0 {
			body.Code = code
		} else {
			body.Code = -1
		}
		body.Msg = err.Error()
		body.Data = resp
	} else {
		body.Msg = "OK"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
func Write(w http.ResponseWriter, res string) {
	//"application/json"
	//w.Header().Set("Content-Type",  "application/json")
	w.WriteHeader(200)
	if n, err := w.Write([]byte(res)); err != nil {
		if err != http.ErrHandlerTimeout {
			logx.Errorf("write response failed, error: %s", err)
		}
	} else if n < len(res) {
		logx.Errorf("actual bytes: %d, written bytes: %d", len(res), n)
	}
}
