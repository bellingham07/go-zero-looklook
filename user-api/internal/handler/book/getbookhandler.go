package book

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-looklook/user-api/internal/logic/book"
	"go-zero-looklook/user-api/internal/svc"
	"go-zero-looklook/user-api/internal/types"
)

func GetBookHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BookReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := book.NewGetBookLogic(r.Context(), svcCtx)
		resp, err := l.GetBook(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
