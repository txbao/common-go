package {{.PkgName}}

import (
	"net/http"
	"encoding/json"
	"github.com/txbao/common-go/common/utils/validators"
	"github.com/txbao/common-go/common/reponse"
	"github.com/txbao/common-go/common/utils"

    "github.com/zeromicro/go-zero/rest/httpx"
	/*{{if .After1_1_10}}"github.com/zeromicro/go-zero/rest/httpx"{{end}}*/
	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.Error(w, err)
			reponse.Response(w, nil, err,0)
			return
		}

		if !(req.Unsigned == "1" && (svcCtx.Config.Mode == "dev" || svcCtx.Config.Mode == "debug")) {
            rq := utils.DesCbcDecryptMap(req.Rq)
            fmt.Println("{{.Call}}参数：",utils.Map2Json(rq))
            if err := json.Unmarshal([]byte(utils.Map2Json(rq)), &req); err != nil {
                reponse.Response(w, nil, err, 0)
                return
            }
        }

        if err := validators.Valid(req); err != nil {
            reponse.Response(w, nil, err, 0)
            return
        }

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx, r)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		//resp, err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		{{if .HasResp}}reponse.Response(w, resp, err,0){{else}}reponse.Response(w, nil, err,0){{end}}
		/*
		if err != nil {
			httpx.Error(w, err)
		} else {
			{{if .HasResp}}httpx.OkJson(w, resp){{else}}httpx.Ok(w){{end}}
		}
		*/
	}
}
