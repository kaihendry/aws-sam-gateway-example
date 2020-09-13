package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	goserverlessapi "github.com/kaihendry/aws-sam-gateway-example"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.Handler("GET", "/healthz", http.HandlerFunc(goserverlessapi.HealthHandler))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router), nil)
}
