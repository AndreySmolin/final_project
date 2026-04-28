package server

import (
	"final_project/pkg/api"
	"log"
	"net/http"
	"time"
)

type Server struct {
	Logger *log.Logger
	Http   *http.Server
}

func NewRouter(logger *log.Logger) *Server {
	router := http.NewServeMux()
	api.Init(router)
	httpServer := &http.Server{
		Addr:         ":7540",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	return &Server{
		Logger: logger,
		Http:   httpServer,
	}
}
