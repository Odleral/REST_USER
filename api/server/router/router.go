package router

import (
	"github.com/go-chi/chi"
	"gitlab.com/odleral/geoportal-go/api/server/requestlog"
	"gitlab.com/odleral/geoportal-go/api/server/router/middleware"
	"gitlab.com/odleral/geoportal-go/cmd/app/app"
)

func New(a *app.App) *chi.Mux {
	l := a.Logger()
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router){
		r.Use(middleware.ContentTypeJson)
		r.Method("POST", "/user/create", requestlog.NewHandler(a.HandleCreateUser, l))
		r.Method("POST", "/user/auth", requestlog.NewHandler(a.HandleAuthUser, l))
		r.Method("PUT", "/user/{uuid}", requestlog.NewHandler(a.HandleUpdateUser, l))
		r.Method("GET", "/user/{uuid}", requestlog.NewHandler(a.HandleFindUser, l))
		r.Method("DELETE", "/user/{uuid}", requestlog.NewHandler(a.HandleDeleteUser, l))
	})

	return r
}
