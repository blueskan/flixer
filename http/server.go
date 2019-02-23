package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/blueskan/flixer/config"
)

type FlixerHttpServer interface {
	Start()
	Stop() error
}

type flixerHttpServer struct {
	server *http.Server
}

func (fhs *flixerHttpServer) Start() {
	go func() {
		fhs.server.ListenAndServe()
	}()
}

func (fhs flixerHttpServer) Stop() error {
	return fhs.server.Shutdown(context.Background())
}

func NewFlixerHttpServer(routes FlixerRoutes, config config.Config) FlixerHttpServer {
	httpServer := &flixerHttpServer{
		server: &http.Server{Addr: fmt.Sprintf(":%s", config.Port)},
	}

	routes.DefineRouteForRenderTemplate(config.Routes.RenderTemplateRoute, config.Template)
	routes.DefineRouteForObtainInputs(config.Routes.ObtainInputRoute)

	return httpServer
}
