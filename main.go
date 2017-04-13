package main

import (
	"log"
	"net/http"

	"github.com/ablease/sleeper-service/sleeperservice"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	sleeperservice.AttachRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
