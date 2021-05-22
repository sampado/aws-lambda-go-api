package main

import (
	"log"
	"net/http"

	"github.com/apex/gateway"
	"github.com/sampado/aws-lambda-go-api/app/hello"
)

func main() {
	http.HandleFunc("/hola", hello.Handler)
	http.HandleFunc("/", hello.Handler)
	addr := ":3000"
	log.Printf("Listening on %s", addr)
	log.Fatal(gateway.ListenAndServe(addr, nil))
}
