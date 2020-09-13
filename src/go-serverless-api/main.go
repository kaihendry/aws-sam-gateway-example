package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/pat"
	goserverlessapi "github.com/kaihendry/aws-sam-gateway-example/src"
)

func main() {
	app := pat.New()
	app.Get("/", goserverlessapi.Get)
	app.Post("/", goserverlessapi.Post)
	app.Post("/healthz", goserverlessapi.HealthHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), app), nil)
}
