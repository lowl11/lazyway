package routes

import "net/http"

type Client struct {
	pattern     string
	hosts       []string
	port        string
	methods     []string
	contentType string
}

func Create(pattern string) *Client {
	return &Client{
		pattern: pattern + "*",
		hosts:   []string{},
		port:    ":80",
		methods: []string{
			http.MethodGet, http.MethodPost,
			http.MethodPut, http.MethodDelete,
		},
		contentType: "application/json",
	}
}
