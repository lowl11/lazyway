package requests

import "net/http"

func (service *Service) setCookies() {
	for key, value := range service.cookies {
		service.request.AddCookie(&http.Cookie{
			Name:  key,
			Value: value,
		})
	}
}

func (service *Service) setHeaders() {
	for key, value := range service.headers {
		for _, item := range value {
			service.request.Header.Add(key, item)
		}
	}
}

func (service *Service) setBasicAuth() {
	if !service.isBasicAuth {
		return
	}

	service.request.SetBasicAuth(service.username, service.password)
}
