package main

import (
	"log"

	"github.com/apex/gateway"
	eg "github.com/kaihendry/aws-sam-gateway-example"
)

func main() {
	log.Fatal(gateway.ListenAndServe("", eg.App()), nil)
}
