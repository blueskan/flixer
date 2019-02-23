package strategy

import (
	"fmt"
)

type StdOutStrategy struct{}

func (sos StdOutStrategy) Process(values string) {
	fmt.Println(values)
}

func NewStdOutStrategy() OutputStrategy {
	return &StdOutStrategy{}
}
