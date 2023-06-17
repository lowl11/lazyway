package interfaces

import "github.com/labstack/echo/v4"

type IRouter interface {
	SetRoute(
		pattern, contentType string,
		hosts []string,
		port string,
		middlewareFunc echo.MiddlewareFunc,
	)
}
