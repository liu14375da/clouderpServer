package handler

import (
	"net/http"

	"ZeroProject/Api/wxUnfolloww/internal/logic"
	"ZeroProject/Api/wxUnfolloww/internal/svc"
	"ZeroProject/Api/wxUnfolloww/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func WxUnfollowwHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewWxUnfollowwLogic(r.Context(), ctx)
		resp, err := l.WxUnfolloww(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
