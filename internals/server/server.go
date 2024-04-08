package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"tools/internals/cfg"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(h http.Handler) error {
	config, err := cfg.LoadAndStore("../cfg")
	if err != nil {
		log.Fatalf("Couldn't parse config: %v", err)
	}
	s.httpServer = &http.Server{
		Addr:           fmt.Sprintf("%s%s", config.Server.Host, config.Server.Port), //все ок?
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
