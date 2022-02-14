package router

import (
	"github.com/go-chi/chi"
	"gitlab.com/odleral/geoportal-go/cmd/app/app"
	"gitlab.com/odleral/geoportal-go/cmd/app/requestlog"
	"gitlab.com/odleral/geoportal-go/cmd/app/router/middleware"
)

func New(a *app.App) *chi.Mux {
	l := a.Logger()
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router){
		r.Use(middleware.ContentTypeJson)
		//User handlers
		r.Method("POST", "/user", requestlog.NewHandler(a.HandleCreateUser, l))
		r.Method("PUT", "/user/{uuid}", requestlog.NewHandler(a.HandleUpdateUser, l))
		r.Method("GET", "/user/{uuid}", requestlog.NewHandler(a.HandleFindUser, l))
		r.Method("DELETE", "/user/{uuid}", requestlog.NewHandler(a.HandleDeleteUser, l))
	})

	//Kubelet checker
	//r.Get("/healthz/liveness", app.HandleLive)
	//r.Method("GET", "/healthz/readiness", requestlog.NewHandler(a.HandleReady, l))

	return r
}
