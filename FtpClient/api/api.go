package api

import "net/http"

func Run_web_gui() {
	http.HandleFunc("/connect", ConnectHandler)
	http.HandleFunc("/list/server", ListServerHandler)
	http.HandleFunc("/list/local", ListLocalHandler)
	http.HandleFunc("/status", StatusHandler)
	http.HandleFunc("/files/upload", UploadFileHandler)
	http.HandleFunc("/files/download", DownloadFileHandler)
	http.HandleFunc("/files/remove", RemoveFileHandler)
	http.HandleFunc("/directories/create", CreateDirectoryHandler)
	http.HandleFunc("/directories/remove", RemoveDirectoryHandler)
	http.HandleFunc("/directories/download", DownloadDirectoryHandler)
	http.HandleFunc("/directories/upload", UploadDirectoryHandler)
	http.ListenAndServe(":5035", nil)
}
