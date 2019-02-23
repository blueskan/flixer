package config

type RouteDefinitions struct {
	RenderTemplateRoute string
	ObtainInputRoute    string
}

type Config struct {
	Routes   RouteDefinitions
	Template string
	Port     string
	Url      string
}
