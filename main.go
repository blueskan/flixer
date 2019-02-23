package main

import (
	"fmt"
	"net/url"

	"github.com/blueskan/flixer/manager/strategy"

	"github.com/blueskan/flixer/config"
	"github.com/blueskan/flixer/http"
	"github.com/blueskan/flixer/manager"
)

func main() {
	config := config.Config{
		Port:     "9000",
		Template: "./template/flixer.html",
		Routes: config.RouteDefinitions{
			ObtainInputRoute:    "/post",
			RenderTemplateRoute: "/",
		},
	}

	obtainInputsCh := make(chan url.Values)

	routes := http.NewFlixerRoutes(obtainInputsCh)
	server := http.NewFlixerHttpServer(routes, config)

	server.Start()

	outputStrategy := strategy.NewStdOutStrategy()
	manager := manager.NewManager(outputStrategy)

	host := "http://localhost"
	manager.OpenInBrowser(fmt.Sprintf("%s:%s%s", host, config.Port, config.Routes.RenderTemplateRoute))

	manager.FinishProcess(obtainInputsCh)
}
