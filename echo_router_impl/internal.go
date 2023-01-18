package echo_router_impl

import (
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazy-gateway/requests"
	"io/ioutil"
	"math/rand"
	"net/http"
)

func (client *Client) handler(ctx echo.Context) error {
	hosts := ctx.Get("gateway_hosts").([]string)
	contentType := ctx.Get("content_type").(string)

	var requestBodyInBytes []byte
	if ctx.Request().Body != nil {
		requestBody, err := ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			return client.Error(ctx, err, "[Gateway] Parse request body error")
		}

		if requestBody != nil {
			requestBodyInBytes = requestBody
		}
	}

	sendHost := hosts[rand.Intn(len(hosts))]
	sendUrl := sendHost + ctx.Request().URL.String()

	response, err := requests.New(ctx.Request().Method, sendUrl, requestBodyInBytes).Send()
	if err != nil {
		return client.Error(ctx, err, "[Gateway] Send request to another host error")
	}

	return ctx.Blob(http.StatusOK, contentType, response)
}
