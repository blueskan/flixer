package strategy

import "fmt"

type StrategyFactory interface {
	GetStrategy(strategy string, filename string) (OutputStrategy, error)
}

type strategyFactory struct{}

func (sf strategyFactory) GetStrategy(
	strategy string,
	filename string,
) (OutputStrategy, error) {
	switch strategy {
	case "stdout":
		return NewStdOutStrategy(), nil
	case "file":
		return NewFileStrategy(filename), nil
	default:
		return nil, fmt.Errorf("Strategy `%s` not available. Try `flixer run --help` for available strategies..", strategy)
	}
}

func NewStrategyFactory() StrategyFactory {
	return &strategyFactory{}
}
