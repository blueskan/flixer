package strategy

import "net/url"

type OutputStrategy interface {
	Process(values url.Values)
}
