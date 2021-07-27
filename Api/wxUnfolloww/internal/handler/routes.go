// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"ZeroProject/Api/wxUnfolloww/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/wx/wx_unfollow",
				Handler: WxUnfollowwHandler(serverCtx),
			},
		},
	)
}
