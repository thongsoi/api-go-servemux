package main

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")

	_, err := DB.Exec("INSERT INTO users2 (name, email) VALUES ($1, $2)", name, email)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/success.html"))
	tmpl.Execute(w, "User saved successfully.")
}
