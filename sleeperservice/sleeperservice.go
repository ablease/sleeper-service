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
	r.HandleFunc("/comms", sleepyMessage).Methods("GET")
}

func sleepyMessage(w http.ResponseWriter, r *http.Request) {
	response := SleepyResponse{
		Level: "INFO",
		Msg:   "you have reached the sleepy service",
		Time:  time.Now(),
	}

	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Errorf("error marshalling response: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(resp), r.URL.Path[1:])
}
