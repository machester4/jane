package lib

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

// I+D
// https://en.wikipedia.org/wiki/BK-tree
// https://en.wikipedia.org/wiki/Edit_distance
// https://blog.algolia.com/inside-the-algolia-engine-part-2-the-indexing-challenge-of-instant-search/

var once sync.Once

// tree bktree, t *trie, tolerance int
func addRecommends(tree bktree, t *Trie, tolerance int) func(*field) {
	return func(w *field) {

		// Skip if the word equals a recommendation
		if w.Same {
			fmt.Println("es igual")
			return
		}

		exact := t.HasKeysWithPrefix(w.Value)
		if exact {
			w.Same = exact
			fmt.Println("es igual")
			return
		}

		// var bkn bktree

		// branch := t.PrefixSearch(w.Value)
		/* for _, candidate := range branch {
			w.Recommends = append(w.Recommends, &RecommendItemSuggestion{
				Entry:    candidate,
				Distance: 0,
			})
			// bkn.add(candidate)
		} */

		// Get results from bk-tree
		results := tree.search(w.Value, tolerance)

		// Sort results from smallest to longest distance
		sort.Slice(results, func(i, j int) bool { return results[i].Distance < results[j].Distance })

		// Add result to word recommends
		for i, result := range results {
			// Skip if the result is equal to the word
			if result.Distance == 0 {
				w.Same = true
				break
			}

			// Skip if we have already reached the maximum of results
			if (i + len(w.Recommends)) == maxResults {
				break
			}
			w.Recommends = append(w.Recommends, result)
		}
	}
}

// Recommend - Contexts are dictionaries with the respective words.
// have higher priority than the words of the language itself
func Recommend(text string, lang string, context string) (*RecommendResult, error) {
	// Create result set
	var rs RecommendResult

	// Create chain
	c := createChain(text)

	// Words provider
	provider, err := getProviderHandler()
	if err != nil {
		return &rs, err
	}

	// Get bk-tree from provider
	cTree, err := provider.getTree(context + "bk")
	if err != nil {
		return &rs, err
	}

	dTree, err := provider.getTree(lang + "bk")
	if err != nil {
		return &rs, err
	}

	cTrie := provider.getTrie(context + "rx")
	dTrie := provider.getTrie(lang + "rx")

	stages := []*stage{
		{
			name:   "Add words recommends from dictionary",
			worker: addRecommends(dTree, dTrie, maxDistanceInDic),
		},
		{
			name:   "Add words recommends from context",
			worker: addRecommends(cTree, cTrie, maxDistanceInContext),
		},
	}

	pipeline := pipeline{stages: stages, fields: c.words}
	pipeline.run()

	// Fill result set
	for _, w := range c.words {
		if w.Recommends != nil && w.Same == false {
			rs.Words = append(rs.Words, w)
		}
	}

	return &rs, nil
}

// Initialize word provider // Indexing
func Initialize(providers ...string) {
	// Create handler instance
	once.Do(func() {
		// Create cache storage for BK-TREES
		provider = &handler{storage: cache.New(5*time.Minute, 10*time.Minute)}

		// Get all word from providers and create BK-TREES
		for _, p := range providers {
			var b bktree
			t := NewTrie()
			for _, w := range getWordsFromFile(p) {
				b.add(w)
				t.Add(w, nil)
			}
			provider.storage.Set(p+"rx", t, cache.NoExpiration)
			provider.storage.Set(p+"bk", b, cache.NoExpiration)
		}
	})
}
