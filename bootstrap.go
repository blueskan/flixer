package main

import (
	"fmt"
	"net/url"

	"github.com/blueskan/flixer/config"
	"github.com/blueskan/flixer/http"
	"github.com/blueskan/flixer/manager"
	"github.com/blueskan/flixer/manager/strategy"
)

func bootstrap(config config.Config) {
	obtainInputsCh := make(chan url.Values)

	routes := http.NewFlixerRoutes(obtainInputsCh)
	server := http.NewFlixerHttpServer(routes, config)

	server.Start()

	outputStrategy := strategy.NewStdOutStrategy()
	manager := manager.NewManager(outputStrategy)

	manager.OpenInBrowser(fmt.Sprintf(
		"%s:%s%s",
		config.Url,
		config.Port,
		config.Routes.RenderTemplateRoute,
	))

	manager.FinishProcess(obtainInputsCh)
}
