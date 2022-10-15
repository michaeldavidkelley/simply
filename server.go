package simply

import (
	"net/http"
)

type Server interface {
	SetAddr(addr string) Server
	SetRouter(router Router) Server
	SetLogger(logger Logger) Server
	Start()
}

type server struct {
	address string
	router  Router
	log     Logger
}

func (s *server) SetLogger(logger Logger) Server {
	s.log = logger
	return s
}

func (s *server) SetRouter(router Router) Server {
	s.router = router

	return s
}

func (s *server) SetAddr(addr string) Server {
	s.address = addr

	return s
}

func (s *server) Start() {
	httpServer := &http.Server{
		Addr:    s.address,
		Handler: s.router,
	}

	err := httpServer.ListenAndServe()

	s.log.Err(err).Error("server failed")
}

func NewServer() Server {
	return &server{}
}

func Start() {
	logger := NewLogger()

	NewServer().
		SetLogger(logger).
		SetAddr(":8080").
		SetRouter(globalRouter.SetLogger(logger)).
		Start()
}
