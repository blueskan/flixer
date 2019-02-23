package strategy

import (
	"fmt"
	"net/url"
)

type StdOutStrategy struct{}

func (sos StdOutStrategy) Process(values url.Values) {
	for key, value := range values {
		fmt.Printf("%s:%s\n", key, value)
	}
}

func NewStdOutStrategy() OutputStrategy {
	return &StdOutStrategy{}
}
