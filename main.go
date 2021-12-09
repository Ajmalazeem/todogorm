package main

import (

	"log"
	"net/http"
	httptransport "github.com/go-kit/kit/transport/http"
)


func main() {
	svc := helloService{}

	helloHandler := httptransport.NewServer(
		makeHelloEndpoint(svc),
		decodeHelloResponse,
		encodeResponse,
	)

	http.Handle("/hello",helloHandler)
	log.Fatal(http.ListenAndServe(":8080",nil))

}