package simply

import (
	"context"
	"net/http"
)

var globalRouter Router = NewRouter()

type Router interface {
	http.Handler
	SetLogger(logger Logger) Router

	Get(path string, c Controller)
}

type router struct {
	logger Logger
	routes map[string]Controller
}

func (r *router) Get(path string, c Controller) {
	r.routes[path] = c
}

func (r *router) SetLogger(logger Logger) Router {
	r.logger = logger

	return r
}

func NewRouter() Router {
	return &router{
		routes: map[string]Controller{},
	}
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	request := NewRequest(req.WithContext(context.WithValue(req.Context(), "log", r.logger)))
	r.logger.With(F{
		"url":    req.URL,
		"query":  req.URL.Query(),
		"method": req.Method,
	}).Info("request")

	var controller Controller
	if c, ok := r.routes[req.RequestURI]; ok {
		controller = c
	} else {
		controller = &notFoundController{}
	}

	r.fromResponse(w, controller.Invoke(request))
}

func WithRouter(fn func(r Router)) {
	fn(globalRouter)
}

func (r *router) fromResponse(w http.ResponseWriter, resp Response) {
	r.logger.With(F{
		"status": resp.GetStatus(),
	}).Info("response")

	w.WriteHeader(resp.GetStatus())

	_, err := w.Write([]byte(resp.GetContents()))
	if err != nil {
		r.logger.Err(err).Error("could not write contents")
	}

}
