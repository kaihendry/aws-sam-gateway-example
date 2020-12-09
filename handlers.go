package awssamgatewayexample

import (
	"fmt"
	"net/http"
	"time"

	"github.com/apex/log"
	"github.com/gorilla/pat"
)

func App() *pat.Router {
	app := pat.New()
	app.Get("/", get)
	app.Post("/", post)
	return app
}

func get(w http.ResponseWriter, r *http.Request) {
	log.Info("saying hello at current time")
	fmt.Fprintf(w, "GET response: hello at %s!\n", time.Now())
}

func post(w http.ResponseWriter, r *http.Request) {
	log.Info("going Postal")
	fmt.Fprintf(w, "POST response: hello at %s!\n", time.Now())
}
