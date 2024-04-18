package main

import (
	"net/http"
)



func main() {
	http.HandleFunc("/connect", connectHandler)
	http.HandleFunc("/list/server", listServerHandler)
	http.HandleFunc("/list/local", listLocalHandler)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/files/upload", uploadHandler)
	http.HandleFunc("/files/download", downloadHandler)
	http.HandleFunc("/files/remove", removeFileHandler)	
	http.HandleFunc("/directories/create", createDirectoryHandler)
	http.HandleFunc("/directories/remove", removeDirectoryHandler)
	http.HandleFunc("/directories/download", downloadDirectoryHandler)
	http.ListenAndServe(":5035", nil)
}
