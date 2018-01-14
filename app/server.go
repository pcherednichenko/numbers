package app

import (
	"log"
	"net/http"
)

const (
	serverStartInfoTpl = "Server start on port: %s"
)

// Server with server mux
type Server struct {
	mux *http.ServeMux
}

// New creates server
func New() *Server {
	return &Server{mux: http.NewServeMux()}
}

// Start server which handle requests
func (s *Server) Start(bind string) error {
	s.mux.HandleFunc("/numbers", numbersHandler)
	log.Printf(serverStartInfoTpl, bind)
	return http.ListenAndServe(bind, s.mux)
}
