package simply

import (
	"context"
	"net/http"
)

type request struct {
	httpRequest *http.Request
}

func (r request) Context() context.Context {
	return r.httpRequest.Context()
}

func (r request) Log() Logger {
	return r.Context().Value("log").(Logger)
}

type Request interface {
	Context() context.Context
	Log() Logger
}

func NewRequest(r *http.Request) Request {
	return &request{httpRequest: r}
}
