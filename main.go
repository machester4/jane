package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/machester4/jane/pipeline"

	"github.com/machester4/jane/lib"
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
		lib.Suggest("!hola ?. estas ahi", "chivito", "es")
		w.Write([]byte(`{"message": "get called"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func main() {
	// http.HandleFunc("/", home)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	block := pipeline.Block{
		Index:    0,
		Value:    "H",
		Category: "letter",
	}
	chain := pipeline.Chain{
		Blocks: []pipeline.Block{block},
	}

	cp := pipeline.ChainPipe{
		Name:    "Log",
		Delayed: false,
		Task: func(chain *pipeline.Chain) {
			fmt.Println("Chain en task", chain)
		},
	}

	bp := pipeline.BlockPipe{
		Name:    "Log block",
		Delayed: false,
		Task: func(block *pipeline.Block) {
			fmt.Println("Chain en task", chain)
		},
	}

	cps := []pipeline.ChainPipe{cp}
	bps := []pipeline.BlockPipe{bp}

	pipe := pipeline.New(&chain, cps, bps)
	fmt.Println("Pipe", pipe)
}
