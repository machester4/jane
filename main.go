package main

import (
	"encoding/json"
	"fmt"
	"github.com/machester4/jane/bktree"
	"github.com/machester4/jane/chain"
	"github.com/machester4/jane/helpers"
	"github.com/machester4/jane/recommender"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
	"os"
	"time"
)

type Body struct {
	Text  string
	Dict  string
	Contexts []string
}

var cach *cache.Cache
func initCache() *cache.Cache {
	// Load tree in mem cache
	if cach == nil {
		fmt.Println("Creando el bk-tree en cache")
		cach = cache.New(5*time.Minute, 10*time.Minute)
		var tree bktree.BKTree
		wordsDict := helpers.GetDictionary("es-ES")
		for _, w := range wordsDict {
			tree.Add(recommender.BkWord(w))
		}
		cach.Set("bk-tree", tree, cache.NoExpiration)
	}
	return cach
}


func draft(w http.ResponseWriter, r *http.Request) {
	// Body struct
	b := Body{}

	// Parse JSON request body
	err := json.NewDecoder(r.Body).Decode(&b)
	helpers.CheckError(err)

	// Marshal or convert recommend object to json and write to response
	c := chain.New(b.Text)

	// BK-TREE
	tree, e := cach.Get("bk-tree")

	fmt.Println(e)

	recommender.Recommend(c, b.Dict, tree.(bktree.BKTree) )
	respJson, err := json.Marshal(c)
	helpers.CheckError(err)

	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write JSON response
	w.Write(respJson)
}

func main() {
	// Init cache
	initCache()

	// Basic server ONLY for test
	http.HandleFunc("/", draft)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
