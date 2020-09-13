package main

import (
	"log"
	"net/http"

	"github.com/apex/gateway"
	goserverlessapi "github.com/kaihendry/aws-sam-gateway-example"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.Handler("GET", "/healthz", http.HandlerFunc(goserverlessapi.HealthHandler))
	log.Fatal(gateway.ListenAndServe("", router), nil)
}
