package main

import (
	"log"

	"github.com/kyosu-1/distributed-system-in-go/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	// ouput starting message
	log.Println("starting server on", srv.Addr)
	// start server
	log.Fatal(srv.ListenAndServe())
}
