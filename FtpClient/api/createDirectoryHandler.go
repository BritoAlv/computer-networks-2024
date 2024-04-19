package api

import (
	"FTPClient/core"
	"encoding/json"
	"net/http"
)


func CreateDirectoryHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	} else if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var request PathRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ftpSession, err := core.SessionBuilder(ftp_to_use)
	if err != nil {
		js, _ := json.Marshal(ResponseOperation{err.Error(), false})
		responseWrite(&w, js)
		return
	}
	stat, err := ftpSession.MKD(request.Path)
	if err != nil {
		js, _ := json.Marshal(ResponseOperation{err.Error(), false})
		responseWrite(&w, js)
		return
	}
	StatusQueue.Enqueue(stat)
	core.SessionFinish(ftpSession)
	
	js, _ := json.Marshal(ResponseOperation{"Directory Created", true})
	responseWrite(&w, js)
}