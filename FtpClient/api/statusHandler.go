package api

import (
	"encoding/json"
	"net/http"
)

type ResponseStatus struct {
	Status string `json:"status"`
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	} else if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if StatusQueue.Len() != 0 {
		js, err := json.Marshal(ResponseStatus{StatusQueue.Dequeue()})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

	if default_session == nil {
		js, err := json.Marshal(ResponseConnect{"Not connected", false})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		return
	}
	stat, err := default_session.NOOP("")
	if err != nil {
		js, err := json.Marshal(ResponseConnect{"Not connected", false})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		return
	}
	StatusQueue.Enqueue(stat)
	js, err := json.Marshal(ResponseStatus{"Connected"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
