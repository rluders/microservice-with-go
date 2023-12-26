package rest

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"menu-service/internal/config"
)

type Server struct {
	*http.Server
	Config *config.ServerHTTP
}

func NewServer(cfg *config.ServerHTTP, router *mux.Router) (*Server, error) {
	if cfg == nil {
		return nil, fmt.Errorf("invalid server config")
	}
	if router == nil {
		return nil, fmt.Errorf("invalid server router")
	}

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	return &Server{
		Server: &http.Server{
			Addr:    addr,
			Handler: router,
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
		},
		Config: cfg,
	}, nil
}

func (s *Server) Start() error {
	var err error

	log.Printf("starting HTTP server at '%s:%d\n", s.Config.Host, s.Config.Port)

	if s.Config.UseHTTPS {
		log.Println("SSL certificate enabled")
		certPath := s.Config.CertPath
		err = s.Server.ListenAndServeTLS(
			fmt.Sprintf("%s/server.crt", certPath),
			fmt.Sprintf("%s/server.key", certPath),
		)
	} else {
		log.Println("SSL certificate disabled")
		err = s.Server.ListenAndServe()
	}

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("unable to start HTTP server: %+v\n", err)
		return err
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) {
	log.Println("HTTP Server shutdown started")
	err := s.Server.Shutdown(ctx)
	if err != nil {
		log.Printf("HTTP Server shutdown failed: %+v\n", err)
		return
	}
	log.Println("HTTP Server shutdown finished")
}
