package main

import (
	"fmt"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	srv := server.NewServer(logger)

	fmt.Println("Starting server on localhost:8080...")
	if err := srv.Server.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}

}
