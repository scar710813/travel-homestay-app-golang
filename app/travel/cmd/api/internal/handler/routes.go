// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	homestay "golodge/app/travel/cmd/api/internal/handler/homestay"
	homestayBussiness "golodge/app/travel/cmd/api/internal/handler/homestayBussiness"
	homestayComment "golodge/app/travel/cmd/api/internal/handler/homestayComment"
	"golodge/app/travel/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/homestay/businessList",
				Handler: homestay.BusinessListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/homestay/guessList",
				Handler: homestay.GuessListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/homestay/homestayDetailWithoutLogin",
				Handler: homestay.HomestayDetailWithoutLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/homestay/homestayList",
				Handler: homestay.HomestayListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/homestay/searchByLocation",
				Handler: homestay.SearchByLocationHandler(serverCtx),
			},
		},
		rest.WithPrefix("/travel/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/homestay/addGuess",
				Handler: homestay.AddGuessHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/homestay/addHomestay",
				Handler: homestay.AddHomestayHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/homestay/adminDeleteHomestay",
				Handler: homestay.AdminDeleteHomestayHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/homestay/deleteHomestay",
				Handler: homestay.DeleteHomestayHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/homestay/homestayDetail",
				Handler: homestay.HomestayDetailHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/travel/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/homestayBussiness/goodBoss",
				Handler: homestayBussiness.GoodBossHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/homestayBussiness/homestayBussinessDetail",
				Handler: homestayBussiness.HomestayBussinessDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/homestayBussiness/homestayBussinessList",
				Handler: homestayBussiness.HomestayBussinessListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/travel/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/homestayComment/commentList",
				Handler: homestayComment.CommentListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/homestayComment/likeComment",
				Handler: homestayComment.LikeCommentHandler(serverCtx),
			},
		},
		rest.WithPrefix("/travel/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/homestayComment/addComment",
				Handler: homestayComment.AddCommentHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/travel/v1"),
	)
}
