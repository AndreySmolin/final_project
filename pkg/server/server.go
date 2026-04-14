package server

import (
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
	fs := http.FileServer(http.Dir("./web"))
	router.Handle("/css/", fs)
	router.Handle("/js/", fs)
	router.Handle("/", fs)
	httpServer := &http.Server{
		Addr:         ":7540",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return &Server{
		Logger: logger,
		Http:   httpServer,
	}
}
