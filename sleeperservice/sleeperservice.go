package sleeperservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	"time"

	"github.com/gorilla/mux"
)

type SleepyResponse struct {
	Level string    `json:"level"`
	Msg   string    `json:"msg"`
	Time  time.Time `json:"time"`
}

func AttachRoutes(r *mux.Router) {
	r.HandleFunc("/comms", SleepyMessage).Methods("GET")
}

func SleepyMessage(w http.ResponseWriter, r *http.Request) {
	var response SleepyResponse
	response = SleepyResponse{
		Level: "INFO",
		Msg:   "you have reached the sleepy service",
		Time:  time.Now(),
	}
	resp, _ := json.Marshal(response)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(resp), r.URL.Path[1:])
}
