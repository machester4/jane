package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/machester4/jane/recommender"

	"github.com/machester4/jane/chain"
)

type Body struct {
	text  string
	dict  string
	rules string
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		var b Body
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "invalid params"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Println(b.text)
		// lib.Suggest("!hola ?. estas ahi", "chivito", "es")
		w.Write([]byte(`{"message": "get called"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func main() {
	// http.HandleFunc("/", home)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	// Creating Chain from text
	s := "holaaaa mundo xD?"
	c := chain.New(s)

	/* c.Walk(func(b *chain.Block) {
		fmt.Printf("Block from walk: %q\n ", b.Value)
	}) */

	// fmt.Println("Get all blocks", len(c.GetAllBlocks()))
	//fmt.Println("Get all words", c.GetWords())

	var contexts = []string{"ch"}
	recommendations := recommender.Recommend(c, "es", contexts)
	for _, rec := range recommendations {
		fmt.Printf("recommendations %q\n", rec.Recommendations)
	}
}
