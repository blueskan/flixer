package main

import (
	"fmt"

	"github.com/blueskan/flixer/manager/strategy"

	"github.com/blueskan/flixer/config"
	"github.com/blueskan/flixer/http"
	"github.com/blueskan/flixer/manager"
)

func bootstrap(
	strategyChoose string,
	filename string,
	config config.Config,
) {
	obtainInputsCh := make(chan string)

	routes := http.NewFlixerRoutes(obtainInputsCh)
	server := http.NewFlixerHttpServer(routes, config)

	server.Start()

	strategyFactory := strategy.NewStrategyFactory()
	outputStrategy, err := strategyFactory.GetStrategy(strategyChoose, filename)
	if err != nil {
		panic(err)
	}

	manager := manager.NewManager(outputStrategy)

	manager.OpenInBrowser(fmt.Sprintf(
		"%s:%s%s",
		config.Url,
		config.Port,
		config.Routes.RenderTemplateRoute,
	))

	manager.FinishProcess(obtainInputsCh)
}
