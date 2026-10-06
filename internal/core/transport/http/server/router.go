package core_http_server

import (
	"fmt"
	"net/http"

	core_http_middleware "github.com/MaximKachkov/ToDooo/internal/core/transport/http/middleware"
)

type APIVersion string

var (
	APIVersion1 = APIVersion("v1")
	APIVersion2 = APIVersion("v2")
	APIVersion3 = APIVersion("v3")
)

type APIVersionRouter struct {
	*http.ServeMux
	apiVersion APIVersion
	Middleware []core_http_middleware.Middleware
}

func NewAPIVersionRouter(apiVersion APIVersion, middleware ...core_http_middleware.Middleware) *APIVersionRouter {
	return &APIVersionRouter{
		ServeMux:   http.NewServeMux(),
		apiVersion: apiVersion,
		Middleware: middleware,
	}
}

func (a *APIVersionRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		a.Handle(pattern, route.WithMiddleware())

	}
}

func (r *APIVersionRouter) WithMiddleware() http.Handler {
	return core_http_middleware.ChainMiddleware(r, r.Middleware...)
}
