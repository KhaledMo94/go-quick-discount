package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)


type Server struct{
	httpServer *http.Server
}

func New(addr string , handler http.Handler) *Server{
	return &Server{
		httpServer: &http.Server{
			Addr: addr,
			Handler: handler,
			ReadTimeout: 10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout: 10 * time.Second,
		},
	}
}

func (s *Server) Start() {
	go func(){
		slog.Info("http server listening", "addr", s.httpServer.Addr)
		if err := s.httpServer.ListenAndServe(); err != nil {
			slog.Error("http server failed", "error", err)
		}
	}()
}

func (s *Server) Shutdown(ctx context.Context) error {
	slog.Info("shutting down http server")
	return s.httpServer.Shutdown(ctx)
}