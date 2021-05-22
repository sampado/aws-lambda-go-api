package main

import (
	"log"
	"net/http"

	"github.com/apex/gateway/v2"
	"github.com/sampado/aws-lambda-go-api/app/hello"
)

func main() {
	http.HandleFunc("/hello", hello.Handler)
	addr := ":8080"
	log.Printf("Listening on %s", addr)
	log.Fatal(gateway.ListenAndServe(addr, nil))
}
