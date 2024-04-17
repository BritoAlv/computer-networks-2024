package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseConnect struct {
	status string
	succesful bool
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func connectHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	js, err := json.Marshal(ResponseConnect{"Connected", true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	fmt.Fprintf(w, "You've hit the POST endpoint!")
}

func handler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Fprintf(w, "Hello, you've hit the API!")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/connect", connectHandler)
	http.ListenAndServe(":5035", nil)
}