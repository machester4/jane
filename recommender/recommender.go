package recommender

import (
	"github.com/machester4/jane/constants"
	"github.com/machester4/jane/provider"
)

func addRecommendsFromDic(w *Word, tree *BKTree) func()  {
	return func() {
		results := tree.Search(w.Value, constants.MaxDistanceInDic)
		for _, result := range results {
			w.Recommends = append(w.Recommends, result.Str)
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
