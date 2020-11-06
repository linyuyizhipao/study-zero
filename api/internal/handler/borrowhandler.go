package handler

import (
	"net/http"

	"black/api/internal/logic"
	"black/api/internal/svc"
	"black/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func borrowHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BorrowReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewBorrowLogic(r.Context(), ctx)
		err := l.Borrow(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
