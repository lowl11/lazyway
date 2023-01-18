package gateway

import (
	"github.com/lowl11/lazy-gateway/interfaces"
)

type Client struct {
	router interfaces.IRouter
}

func Create(router interfaces.IRouter) *Client {
	return &Client{
		router: router,
	}
}
