package main

import (
	"encoding/json"
	"github.com/machester4/jane/chain"
	"github.com/machester4/jane/helpers"
	"github.com/machester4/jane/recommender"
	"log"
	"net/http"
	"os"
)

type Body struct {
	Text  string
	Dict  string
	Contexts []string
}

func draft(w http.ResponseWriter, r *http.Request) {
	// Body struct
	b := Body{}

	// Parse JSON request body
	err := json.NewDecoder(r.Body).Decode(&b)
	helpers.CheckError(err)

	// Marshal or convert recommend object to json and write to response
	c := chain.New(b.Text)
	recommender.Recommend(c, b.Dict, b.Contexts)
	respJson, err := json.Marshal(c)
	helpers.CheckError(err)

	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write JSON response
	w.Write(respJson)
}

func main() {
	// Basic server ONLY for test
	http.HandleFunc("/", draft)
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
}
