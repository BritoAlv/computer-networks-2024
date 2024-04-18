package api

import (
	"encoding/json"
	"net/http"
	"FTPClient/core"
)

type ListRequest struct {
	Path string `json:"path"`
}

type ListResponse struct {
	Directories []string `json:"directories"`
	Files       []string `json:"files"`
	Successful   bool     `json:"successful"`
}

func ListServerHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	} else if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var request ListRequest
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
	folders, files, err := core.Get_files_folders_Server(ftpSession, request.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	core.SessionFinish(ftpSession)
	js, err := json.Marshal(ListResponse{folders, files, true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}