package simply

import "net/http"

type Controller interface {
	Invoke(req Request) Response
}

type notFoundController struct {
}

func (c *notFoundController) Invoke(req Request) Response {
	return NewResponse().SetStatus(http.StatusNotFound)
}
