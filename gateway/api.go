package gateway

import (
	"github.com/lowl11/lazy-gateway/routes"
)

func (client *Client) Route(route *routes.Client) {
	client.router.SetRoute(
		route.Pattern(),
		route.ContentType(),
		route.Hosts(),
		route.Port(),
		route.Middleware(),
	)
}
