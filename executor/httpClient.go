package executor

import "net/http"

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type httpClient struct {
	client *http.Client
}

func NewHttpClient(client *http.Client) HttpClient {
	return &httpClient{
		client: client,
	}
}

func (httpClient *httpClient) Do(req *http.Request) (*http.Response, error) {
	return httpClient.client.Do(req)
}
