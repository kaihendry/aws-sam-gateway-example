package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/kaihendry/aws-sam-gateway-example"

	"github.com/julienschmidt/httprouter"
)

const (
	serverPort = 8000
)

func main() {
	router := httprouter.New()
	router.Handler("GET", "/healthz", http.HandlerFunc(goserverlessapi.HealthHandler))

	fmt.Printf("Server listening on port: %d\n", serverPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", serverPort), router), nil)
}
