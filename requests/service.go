package requests

import "net/http"

type Service struct {
	method string
	url    string
	body   any

	headers map[string]string
	cookies map[string]string

	isBasicAuth bool
	username    string
	password    string

	request *http.Request

	status   int
	response []byte
}

func New(method, url string, body any) *Service {
	return &Service{
		method: method,
		url:    url,
		body:   body,

		headers: make(map[string]string),
		cookies: make(map[string]string),
	}
}
