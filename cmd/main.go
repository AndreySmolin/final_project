package main

import (
	"final_project/pkg/db"
	"final_project/pkg/server"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "Server:", log.LstdFlags)
	err := db.Init("./internal/db/scheduler.db")
	if err != nil {
		logger.Fatalf("Fatal error:%v", err)
	}
	defer db.CloseDB()
	server := server.NewRouter(logger)
	logger.Print("Start port:7540")
	err = server.Http.ListenAndServe()
	if err != nil {
		logger.Fatalf("Fatal error:%v", err)
	}
}
