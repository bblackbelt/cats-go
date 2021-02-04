package executor

type httpExecutor struct {
	Client HttpClient
}

func NewHttpExecutor(httpClient HttpClient) Executor {
	return &httpExecutor{
		Client: httpClient,
	}
}

func (httpExecutor *httpExecutor) Execute(req *Request) (*Response, error) {
	resp, err := httpExecutor.Client.Do(req.Req)
	if err != nil {
		return nil, err
	}
	return &Response{
		Body: resp.Body,
	}, nil
}
