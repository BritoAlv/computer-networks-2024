package main


import (
	"encoding/json"
	"net/http"
)

type FileTransferRequest struct {
	Source string `json:"source"`
	Destination string  `json:"destination"`
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
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
	ftpSession, err := SessionBuilder(ftp_to_use)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	stat, err := ftpSession.PUT(request.Source + "&" + request.Destination + "/" + get_filename_path(request.Source))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	SessionFinish(ftpSession)
	StatusQueue.Enqueue(stat)
	js, err := json.Marshal(ResponseConnect{"File uploaded", true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}