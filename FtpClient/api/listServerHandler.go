package api

import (
	"FTPClient/core"
	"encoding/json"
	"net/http"
)

func ListServerHandler(w http.ResponseWriter, r *http.Request) {
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
		js, _ := json.Marshal(ListResponse{make([]string, 0),make([]string, 0), false})
		responseWrite(&w, js)
		return
	}
	folders, files, err := core.Get_files_folders_Server(ftpSession, request.Path)
	if err != nil {
		js, _ := json.Marshal(ListResponse{folders, files, false})
		responseWrite(&w, js)
		return
	}
	core.SessionFinish(ftpSession)
	js, _ := json.Marshal(ListResponse{folders, files, true})
	responseWrite(&w, js)
}