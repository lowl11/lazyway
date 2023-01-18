package echo_router_impl

import (
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazy-gateway/echo_router_impl/controller"
	"github.com/lowl11/lazy-gateway/interfaces"
)

type Client struct {
	controller.Base
	server *echo.Echo
}

func Create(server *echo.Echo) interfaces.IRouter {
	return &Client{
		server: server,
	}
}
