package lib

import (
	"sort"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

var once sync.Once

func addRecommends(w *Field, tree bktree, tolerance int) func() {
	return func() {
		// Skip if have all recommends
		if len(w.Recommends) == maxResults {
			return
		}

		// Get results from bk-tree
		results := tree.search(bkword(w.Value), tolerance)

		// Sort results from smallest to longest distance
		sort.Slice(results, func(i, j int) bool { return results[i].Distance < results[j].Distance })

		// Add result to word recommends
		for i, result := range results {
			// Skip if the result is equal to the word
			if i == maxResults || result.Distance == 0 {
				break
			}
			w.Recommends = append(w.Recommends, result)
		}
	}
}

// Contexts are dictionaries with the respective words.
// have higher priority than the words of the language itself
func Recommend(text string, lang string, context string) *Chain {
	// Create chain
	c := createChain(text)

	// Words provider
	provider := getProviderHandler()

	// Get BK-TREE from provider
	cTree := provider.getTree(context)
	dTree := provider.getTree(lang)

	var sContext []func()
	var sDict []func()
	for _, word := range c.Words {
		// Create steps for 'Add words recommends from context' stage
		sContext = append(sContext, addRecommends(word, cTree, maxDistanceInContext))

		// Create steps for 'Add words recommends from dictionary' stage
		sDict = append(sDict, addRecommends(word, dTree, maxDistanceInDic))
	}

	// TODO Create steps for 'Remove forbidden words' stage (use bk-tree for remove forbidden words with distance 1)

	// Create stages for pipeline
	stages := []*stage{
		{
			name:  "Add words recommends from context",
			steps: sContext,
		},
		{
			name:  "Add words recommends from dictionary",
			steps: sDict,
		},
		{
			name:  "Remove forbidden words",
			steps: nil,
		},
	}

	// Start pipeline async
	pipeline := pipeline{stages: stages}
	pipeline.start(false)
	return c
}

func Initialize(providers ...string) {
	// Create handler instance
	once.Do(func() {
		// Create cache storage for BK-TREES
		provider = &handler{storage: cache.New(5*time.Minute, 10*time.Minute)}

		// Get all word from providers and create BK-TREES
		for _, p := range providers {
			var b bktree
			for _, w := range getWordsFromFile(p) {
				b.add(bkword(w))
			}
			provider.storage.Set(p, b, cache.NoExpiration)
		}
	})
}
