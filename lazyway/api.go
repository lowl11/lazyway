package lazyway

import (
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazyway/echo_router_impl"
	"github.com/lowl11/lazyway/gateway"
	"github.com/lowl11/lazyway/routes"
)

func Echo(server *echo.Echo) *gateway.Client {
	return gateway.Create(echo_router_impl.Create(server))
}

func Route(pattern string) *routes.Client {
	return routes.Create(pattern)
}
