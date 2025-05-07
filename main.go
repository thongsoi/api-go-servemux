package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	InitDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/submit", SubmitHandler)

	port := os.Getenv("PORT")
	log.Println("Server listening on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
