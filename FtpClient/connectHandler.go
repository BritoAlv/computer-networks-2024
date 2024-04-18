package main

import (
	"encoding/json"
	"net/http"
)

type ConnectRequest struct {
	IpAddress string `json:"ipAddress"`
	UserName  string `json:"userName"`
	Password  string `json:"password"`
	Port      string `json:"port"`
}

var ftp_to_use FTPExample
var default_session *FtpSession

type ResponseConnect struct {
	Status    string `json:"status"`
	Succesful bool   `json:"successful"`
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

	default_session, err  = SessionBuilder(ftp_to_use)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(ResponseConnect{"OK", true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}