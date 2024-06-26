package client

import (
	"net/http"
)

type HttpClient struct {
    client *http.Client
}

func NewHttpClient() *HttpClient {
    return &HttpClient{
        client: &http.Client{},
    }
}

func (c *HttpClient) Do(req *http.Request) (*http.Response, error) {
    return c.client.Do(req)
}
