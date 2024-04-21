package api

import (
	"FTPClient/core"
	"encoding/json"
	"net/http"
)

func CloseHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	} else if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	core.SessionFinish(default_session)
	js, _ := json.Marshal(ResponseOperation{"Disconnected", true})
	responseWrite(&w, js)
}