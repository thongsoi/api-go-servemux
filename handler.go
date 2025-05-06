package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func APIMessageHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello from Go API!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
