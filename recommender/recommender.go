package recommender

import (
	"fmt"
	"github.com/machester4/jane/bktree"
	"github.com/machester4/jane/constants"
	"github.com/machester4/jane/provider"
)

func addRecommendsFromDic(w *Word, tree bktree.BKTree) func()  {
	return func() {
		var results []*bktree.Result
		for i,rLen := 0, len(results); rLen < constants.MaxResults; i++ {
			fmt.Println("vuelta")
			// Search from smallest to longest distance to obtain the maximum number of results
			tResults := tree.Search(bktree.Word(w.Value), i, constants.MaxResults - rLen)

			// Skip if have result with distance 0
			if i == 0 && len(tResults) > 0 {
				break
			}
			// Skip if the current distance is greater than required
			if i > constants.MaxDistanceInDic {
				break
			}

			// Add results
			results = append(results, tResults...)
		}
		// Add recommends to word
		w.Recommends = results
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
