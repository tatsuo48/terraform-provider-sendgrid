package client

import (
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
)

type SendgridCLient struct {
	APIKey string
	Host   string
}

func (s SendgridCLient) Get(path string) (*rest.Response, error) {
	request := sendgrid.GetRequest(s.APIKey, path, s.Host)
	request.Method = "GET"
	r, err := sendgrid.API(request)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s SendgridCLient) Post(path, body string) (*rest.Response, error) {
	request := sendgrid.GetRequest(s.APIKey, path, s.Host)
	request.Method = "POST"
	request.Body = []byte(body)
	r, err := sendgrid.API(request)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s SendgridCLient) Delete(path string) (*rest.Response, error) {
	request := sendgrid.GetRequest(s.APIKey, path, s.Host)
	request.Method = "DELETE"
	r, err := sendgrid.API(request)
	if err != nil {
		return nil, err
	}
	return r, nil
}
