package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var ftp_to_use FTPExample

type ResponseConnect struct {
	Status    string `json:"status"`
	Succesful bool   `json:"successful"`
}

type ListRequest struct {
	Path string `json:"path"`
}

type ListResponse struct {
	Directories []string
	Files       []string
	Succesful bool
}

type ConnectRequest struct {
	IpAddress string `json:"ipAddress"`
	UserName  string `json:"userName"`
	Password  string `json:"password"`
	Port      string `json:"port"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func connectHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	} else if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var request ConnectRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ftp_to_use = FTPExample{
		ip:       request.IpAddress,
		port:     request.Port,
		user:     request.UserName,
		password: request.Password,
	}

	js, err := json.Marshal(ResponseConnect{"OK", true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func listServerHandler(w http.ResponseWriter, r *http.Request) {
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
	ftpSession, err := SessionBuilder(ftp_to_use)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(request.Path)
	folders, files, err := get_files_folders_current(ftpSession, request.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	SessionFinish(ftpSession)
	js, err := json.Marshal(ListResponse{folders, files, true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/connect", connectHandler)
	http.HandleFunc("/list/server", listServerHandler)
	http.HandleFunc("/list/local", listServerHandler)
	http.ListenAndServe(":5035", nil)
}