package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	addr       string
}

func (s *Server) Run(h http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":8080", //потом передавать сюда нужный порт
		Handler:        h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s.httpServer.ListenAndServe()
}
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
