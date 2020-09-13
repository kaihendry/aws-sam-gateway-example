package main

import (
	"log"

	"github.com/apex/gateway"
	"github.com/gorilla/pat"
	goserverlessapi "github.com/kaihendry/aws-sam-gateway-example/src"
)

func main() {
	app := pat.New()
	app.Get("/", goserverlessapi.Get)
	app.Post("/", goserverlessapi.Post)
	app.Post("/healthz", goserverlessapi.HealthHandler)
	log.Fatal(gateway.ListenAndServe("", app), nil)
}
