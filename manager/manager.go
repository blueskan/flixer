package manager

import (
	"net/url"

	"github.com/blueskan/flixer/manager/strategy"
	"github.com/skratchdot/open-golang/open"
)

type Manager interface {
	OpenInBrowser(url string)
	FinishProcess(obtainInputsCh <-chan url.Values)
}

type manager struct {
	strategy strategy.OutputStrategy
}

func (m manager) OpenInBrowser(url string) {
	open.Run(url)
}

func (m manager) FinishProcess(obtainInputsCh <-chan url.Values) {
	inputs := <-obtainInputsCh

	m.strategy.Process(inputs)
}

func NewManager(strategy strategy.OutputStrategy) Manager {
	return manager{
		strategy: strategy,
	}
}
