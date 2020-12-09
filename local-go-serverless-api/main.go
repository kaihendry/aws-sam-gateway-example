package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	eg "github.com/kaihendry/aws-sam-gateway-example"
)

func main() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), eg.App()), nil)
}
