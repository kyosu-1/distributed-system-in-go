package main

import (
	"log"

	"github.com/kyosu-1/distributed-system-in-go/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
