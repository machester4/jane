package recommender

import (
	"github.com/machester4/jane/bktree"
	"github.com/machester4/jane/constants"
	"github.com/machester4/jane/provider"
	"sort"
)

func addRecommendsFromDic(w *Word, tree bktree.BKTree) func()  {
	return func() {
		// Get results from bk-tree
		results := tree.Search(bktree.Word(w.Value), constants.MaxDistanceInDic)

		// Sort results from smallest to longest distance
		sort.Slice(results, func(i, j int) bool { return results[i].Distance < results[j].Distance })

		// Add result to word recommends
		for i, result := range results {
			// Skip if the result is equal to the word
			if i == constants.MaxResults || result.Distance == 0 {
				break
			}
			w.Recommends = append(w.Recommends, result)
		}
	}
}

// Contexts are dictionaries with the respective words.
// have higher priority than the words of the language itself
func Recommend(chain *Chain, lang string) {
	// Words provider
	provider := provider.GetHandler()

	// Get BK-TREE from provider
	tree := provider.GetTree(lang)

	// Words in chain
	words := chain.Words

	// TODO Create steps for 'Add words recommends from context' stage

	// Create steps for 'Add words recommends from dictionary' stage
	var steps []func()
	for _, word := range words {
		steps = append(steps, addRecommendsFromDic(word, tree))
	}

	// TODO Create steps for 'Remove forbidden words' stage (use bk-tree for remove forbidden words with distance 1)

	// Create stages for pipeline
	stages := []*Stage{
		{
			Name:  "Add words recommends from context",
			Steps: nil,
		},
		{
			Name:  "Add words recommends from dictionary",
			Steps: steps,
		},
		{
			Name:  "Remove forbidden words",
			Steps: nil,
		},
	}

	// Start pipeline async
	pipeline := Pipeline{Stages: stages}
	pipeline.Start(false)
}
