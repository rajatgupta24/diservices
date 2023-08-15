package main

import (
	"log"

	"github.com/rajatgupta24/diservices/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
