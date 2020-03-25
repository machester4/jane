package lib

import (
	"sort"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

// I+D
// https://en.wikipedia.org/wiki/BK-tree
// https://en.wikipedia.org/wiki/Edit_distance

var once sync.Once

func addRecommends(tree bktree, tolerance int) func(*field) {
	return func(w *field) {
		// Skip if the word equals a recommendation
		if w.Same == true {
			return
		}

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
	cTree, err := provider.getTree(context)
	if err != nil {
		return &rs, err
	}

	dTree, err := provider.getTree(lang)
	if err != nil {
		return &rs, err
	}

	stages := []*stage{
		{
			name:   "Add words recommends from dictionary",
			worker: addRecommends(dTree, maxDistanceInDic),
		},
		{
			name:   "Add words recommends from context",
			worker: addRecommends(cTree, maxDistanceInContext),
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

// Initialize word provider
func Initialize(providers ...string) {
	// Create handler instance
	once.Do(func() {
		// Create cache storage for BK-TREES
		provider = &handler{storage: cache.New(5*time.Minute, 10*time.Minute)}

		// Get all word from providers and create BK-TREES
		for _, p := range providers {
			var b bktree
			for _, w := range getWordsFromFile(p) {
				b.add(w)
			}
			provider.storage.Set(p, b, cache.NoExpiration)
		}
	})
}
