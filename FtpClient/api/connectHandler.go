package api

import (
	"FTPClient/core"
	"encoding/json"
	"net/http"
)



var ftp_to_use core.FTPExample
var default_session *core.FtpSession

var StatusQueue core.Queue = *core.NewQueue()

func ConnectHandler(w http.ResponseWriter, r *http.Request) {
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
	ftp_to_use = core.FTPExample{
		Ip:       request.IpAddress,
		Port:     request.Port,
		User:     request.UserName,
		Password: request.Password,
	}

	default_session, err  = core.SessionBuilder(ftp_to_use)
	if err != nil {
		js, _ := json.Marshal(ResponseOperation{err.Error(), false})
		responseWrite(&w, js)
		return
	}

	js, err := json.Marshal(ResponseOperation{"OK", true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	responseWrite(&w, js)
}