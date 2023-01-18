package lazyway

import (
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazy-gateway/echo_router_impl"
	"github.com/lowl11/lazy-gateway/gateway"
	"github.com/lowl11/lazy-gateway/routes"
)

func Echo(server *echo.Echo) *gateway.Client {
	return gateway.Create(echo_router_impl.Create(server))
}

func Route(pattern string) *routes.Client {
	return routes.Create(pattern)
}
