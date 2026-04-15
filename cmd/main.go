package main

import (
	"final_project/pkg/db"
	"final_project/pkg/server"
	"log"
	"os"
)

func main() {
	err := db.Init("./internal/db/scheduler.db")
	logger := log.New(os.Stdout, "Server:", log.LstdFlags)
	server := server.NewRouter(logger)
	logger.Print("Start port:7540")
	if err != nil {
		logger.Fatalf("Fatal error:%v", err)
	}
	err = server.Http.ListenAndServe()
	if err != nil {
		logger.Fatalf("Fatal error:%v", err)
	}
}
