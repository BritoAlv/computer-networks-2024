package api

import (
	"encoding/json"
	"net/http"
)

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
		js, _ := json.Marshal(ResponseStatus{StatusQueue.Dequeue()})
		responseWrite(&w, js)
	}

	if default_session == nil {
		js, _ := json.Marshal(ResponseStatus{"Not connected"})
		responseWrite(&w, js)
		return
	}
	stat, err := default_session.NOOP("")
	if err != nil {
		js, _ := json.Marshal(ResponseStatus{err.Error()})
		responseWrite(&w, js)
		return
	}
	StatusQueue.Enqueue(stat)
	js, _ := json.Marshal(ResponseStatus{"Connected"})
	responseWrite(&w, js)
}
