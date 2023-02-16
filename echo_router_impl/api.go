package echo_router_impl

import (
	"github.com/labstack/echo/v4"
)

func (client *Client) SetRoute(pattern, contentType string, hosts []string, port string, middlewareFunc echo.MiddlewareFunc) {
	group := client.server.Group(pattern)

	// custom middleware
	if middlewareFunc != nil {
		group.Use(middlewareFunc)
	}

	group.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Set("gateway_hosts", hosts)
			ctx.Set("gateway_port", port)
			ctx.Set("content_type", contentType)

			if err := next(ctx); err != nil {
				ctx.Error(err)
				return nil
			}

			return nil
		}
	})

	group.Any("/*", client.handler)
}
