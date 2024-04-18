package main

import (
	"FTPClient/api"
	"net/http"
)

func run_web_gui() {
	http.HandleFunc("/connect", api.ConnectHandler)
	http.HandleFunc("/list/server", api.ListServerHandler)
	http.HandleFunc("/list/local", api.ListLocalHandler)
	http.HandleFunc("/status", api.StatusHandler)
	http.HandleFunc("/files/upload", api.UploadFileHandler)
	http.HandleFunc("/files/download", api.DownloadFileHandler)
	http.HandleFunc("/files/remove", api.RemoveFileHandler)
	http.HandleFunc("/directories/create", api.CreateDirectoryHandler)
	http.HandleFunc("/directories/remove", api.RemoveDirectoryHandler)
	http.HandleFunc("/directories/download", api.DownloadDirectoryHandler)
	http.HandleFunc("/directories/upload", api.UploadDirectoryHandler)
	http.ListenAndServe(":5035", nil)
}

func main() {
	run_web_gui()
}
