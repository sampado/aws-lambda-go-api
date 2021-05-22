package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sampado/aws-lambda-go-api/app/hello"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", hello.Handler)

	addr := ":8080"
	log.Printf("Listening on %s", addr)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatal(err)
	}
}
