package api

import "net/http"

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func responseWrite (w *http.ResponseWriter, js []byte) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Write(js)
}