package main

import (
	"embed"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/apex/log"
	jsonhandler "github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/text"

	"github.com/apex/gateway/v2"
)

//go:embed templates
var tmpl embed.FS
var Version string

func main() {
	t, err := template.ParseFS(tmpl, "templates/*.html")
	if err != nil {
		log.WithError(err).Fatal("Failed to parse templates")
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "text/html")
		err = t.ExecuteTemplate(rw, "index.html", struct {
			Version string
			Time    time.Time
			Header  http.Header
		}{Version, time.Now(), r.Header})
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			log.WithError(err).Fatal("Failed to execute templates")
		}
		val, ok := r.Header["X-Amzn-Trace-Id"]
		if ok {
			log.WithField("trace", val).Info("served request")
		}
	})

	port := os.Getenv("_LAMBDA_SERVER_PORT")
	if port == "" {
		log.SetHandler(text.Default)
		err = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	} else {
		log.SetHandler(jsonhandler.Default)
		err = gateway.ListenAndServe("", nil)
	}
	log.Fatalf("failed to start server: %v", err)
}
