package xor_http_server

import (
	"context"
	"github.com/pkg/errors"
	"net/http"
)

type Server struct {
	srv *http.Server
}

func NewServer(cfg *Config, r *Router) *Server {
	return &Server{
		srv: &http.Server{
			Addr:         ":" + cfg.Port,
			Handler:      r.router,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
		},
	}
}

// TODO support http server with tls

func (s *Server) Start() error {
	if s.srv.Handler == nil {
		return errors.New("no routes have registered")
	}
	if err := s.srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
