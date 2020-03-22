package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/machester4/jane/lib"
)

type Body struct {
	Text    string
	Lang    string
	Context string
}

func draft(w http.ResponseWriter, r *http.Request) {
	// Body struct
	b := Body{}

	// Parse JSON request body
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		panic(err)
	}

	c := lib.Recommend(b.Text, b.Lang, b.Context)

	respJSON, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write JSON response
	w.Write(respJSON)
}

func main() {
	// Initialize words provider handler
	lib.Initialize("es-50", "chivito")

	// Basic server ONLY for test
	http.HandleFunc("/", draft)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
