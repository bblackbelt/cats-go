package executor

import (
	"io"
	"net/http"
)

type Request struct {
	Req *http.Request
}

type Response struct {
	Body io.ReadCloser
}

type Executor interface {
	Execute(req *Request) (*Response, error)
}
