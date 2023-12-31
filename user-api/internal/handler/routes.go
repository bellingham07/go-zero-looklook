// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	book "go-zero-looklook/user-api/internal/handler/book"
	user "go-zero-looklook/user-api/internal/handler/user"
	"go-zero-looklook/user-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/test",
				Handler: user.UserTestHandler(serverCtx),
			},
		},
		rest.WithPrefix("/userapi/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.TestMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/login",
					Handler: user.LoginHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/userapi/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/get_book",
				Handler: book.GetBookHandler(serverCtx),
			},
		},
		rest.WithPrefix("/userapi/v1"),
	)
}
