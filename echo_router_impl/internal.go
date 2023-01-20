package echo_router_impl

import (
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazy-gateway/requests"
	"io/ioutil"
	"math/rand"
	"net/http"
)

func (client *Client) handler(ctx echo.Context) error {
	request := ctx.Request()

	// get need data for certain route
	hosts := ctx.Get("gateway_hosts").([]string)
	contentType := ctx.Get("content_type").(string)

	// read request body is exist
	var requestBodyInBytes []byte
	if request.Body != nil {
		requestBody, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return client.Error(ctx, err, "[Gateway] Parse request body error")
		}

		if requestBody != nil {
			requestBodyInBytes = requestBody
		}
	}

	// headers (with cookies)
	headers := make(map[string]string)
	for key, value := range request.Header {
		if len(value) > 0 {
			headers[key] = value[0]
		}
	}

	// choose need host
	sendHost := hosts[rand.Intn(len(hosts))]
	sendUrl := sendHost + request.URL.String()

	// send request
	response, err := requests.New(request.Method, sendUrl, requestBodyInBytes).
		Headers(headers).
		Send()
	if err != nil {
		return client.Error(ctx, err, "[Gateway] Send request to another host error")
	}

	return ctx.Blob(http.StatusOK, contentType, response)
}
