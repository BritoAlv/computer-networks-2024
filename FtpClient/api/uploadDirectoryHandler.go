package api

import (
	"encoding/json"
	"net/http"
	"FTPClient/core"
)

func UploadDirectoryHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	} else if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var request FileTransferRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ftpSession, err := core.SessionBuilder(ftp_to_use)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	stat, err := ftpSession.RPUT(request.Source + "&" + request.Destination)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	core.SessionFinish(ftpSession)
	StatusQueue.Enqueue(stat)
	
	js, err := json.Marshal(ResponseConnect{"Directory Uploaded", true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}