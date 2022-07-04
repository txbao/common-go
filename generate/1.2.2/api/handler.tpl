package handler

import (
	"net/http"
	"encoding/json"
	"github.com/txbao/common-go/common/reponse"
    "github.com/txbao/common-go/common/utils"
    "fmt"

	{{.ImportPackages}}
	"github.com/zeromicro/go-zero/rest/httpx"
)

func {{.HandlerName}}(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.Error(w, err)
			reponse.Response(w, nil, err,0)
			return
		}{{end}}

		if (req.Unsigned == "1" && (ctx.Config.Mode == "dev" || ctx.Config.Mode == "debug")) == false {
            rq := utils.DesCbcDecryptMap(req.Rq)
            fmt.Println("{{.Call}}参数：",utils.Map2Json(rq))
            json.Unmarshal([]byte(utils.Map2Json(rq)), &req)
        }

		l := logic.New{{.LogicType}}(r.Context(), ctx, r)
		resp, err := l.{{.Call}}({{if .HasRequest}}req{{end}})
		reponse.Response(w, resp, err,0)
		/*
		if err != nil {
			httpx.Error(w, err)
		} else {
			{{if .HasResp}}httpx.OkJson(w, resp){{else}}httpx.Ok(w){{end}}
		}
		*/
	}
}
