package routes

import (
	"github.com/labstack/echo/v4"
	"strings"
)

func (client *Client) Hosts() []string {
	return client.hosts
}

func (client *Client) SetHosts(hosts ...string) *Client {
	client.hosts = hosts
	return client
}

func (client *Client) Port() string {
	return client.port
}

func (client *Client) SetMiddleware(middlewareFunc echo.MiddlewareFunc) *Client {
	client.middlewareFunc = middlewareFunc
	return client
}

func (client *Client) Middleware() echo.MiddlewareFunc {
	return client.middlewareFunc
}

func (client *Client) SetPort(port string) *Client {
	if port == "" || !strings.Contains(port, ":") {
		return client
	}

	client.port = port
	return client
}

func (client *Client) Pattern() string {
	return client.pattern
}

func (client *Client) SetContentType(contentType string) *Client {
	client.contentType = contentType
	return client
}

func (client *Client) ContentType() string {
	return client.contentType
}
