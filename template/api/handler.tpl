package {{.PkgName}}

import (
	"hibug-tuning-platfrom/api/internal/common/httpc"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
            httpc.RespError(r.Context(), w, r, err)
			return
		}
		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		if err != nil {
			httpc.RespError(r.Context(), w, r, err)
		} else {
			{{if .HasResp}}httpc.RespSuccess(r.Context(), w, resp){{else}}httpc.RespSuccess(r.Context(), w, nil){{end}}
		}
	}
}
