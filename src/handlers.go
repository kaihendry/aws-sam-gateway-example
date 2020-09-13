package goserverlessapi

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/apex/log"
)

// pets database-ish
var pets = make(map[string]struct{})

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func Get(w http.ResponseWriter, r *http.Request) {
	log.Info("list pets")

	if len(pets) == 0 {
		fmt.Fprintf(w, "no pets")
		return
	}

	for name := range pets {
		fmt.Fprintf(w, "- %s\n", name)
	}
}

// curl -d Tobi http://localhost:3000/
// curl -d Loki http://localhost:3000/
// curl -d Jane http://localhost:3000/
func Post(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	name := string(b)
	pets[name] = struct{}{}
	log.WithField("name", name).Info("add pet")
	fmt.Fprintf(w, "welcome to the family %s!\n", name)
}
