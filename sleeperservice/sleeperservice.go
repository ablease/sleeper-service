package sleeperservice

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func AttachRoutes(r *mux.Router) {
	r.HandleFunc("/comms", SleepyMessage).Methods("GET")
}

func SleepyMessage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "a sleepy message %s", r.URL.Path[1:])
}
