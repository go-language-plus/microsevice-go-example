package main

import (
	"log"
	"net/http"

	"stringsvc/endpoint"
	"stringsvc/service"
	"stringsvc/transport"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	svc := service.StringSvc{}

	uppercaseHandler := httptransport.NewServer(
		endpoint.MakeUppercaseEndpoint(svc),
		transport.DecodeUppercaseRequest,
		transport.EncodeResponse,
	)

	countHandler := httptransport.NewServer(
		endpoint.MakeCountEndpoint(svc),
		transport.DecodeCountRequest,
		transport.EncodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
