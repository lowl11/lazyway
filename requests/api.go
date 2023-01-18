package requests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (service *Service) Cookies(cookies map[string]string) *Service {
	service.cookies = cookies
	return service
}

func (service *Service) Cookie(key, value string) *Service {
	service.cookies[key] = value
	return service
}

func (service *Service) Headers(headers map[string]string) *Service {
	service.headers = headers
	return service
}

func (service *Service) Header(key, value string) *Service {
	service.headers[key] = value
	return service
}

func (service *Service) Send() ([]byte, error) {
	var buffer *bytes.Buffer
	if service.body != nil {
		if _, ok := service.body.([]byte); ok {
			buffer = bytes.NewBuffer(service.body.([]byte))
		} else {
			bodyInBytes, err := json.Marshal(service.body)
			if err != nil {
				return nil, err
			}

			buffer = bytes.NewBuffer(bodyInBytes)
		}
	}

	request, err := http.NewRequest(service.method, service.url, buffer)
	if err != nil {
		return nil, err
	}

	service.request = request

	service.setHeaders()
	service.setCookies()
	service.setBasicAuth()

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = response.Body.Close(); err != nil {
			log.Println("Close gateway request error")
		}
	}()

	responseInBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	service.status = response.StatusCode
	service.response = responseInBytes

	return responseInBytes, nil
}

func (service *Service) SendStatus() ([]byte, int, error) {
	response, err := service.Send()
	return response, service.status, err
}
