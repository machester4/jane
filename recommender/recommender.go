package recommender

import (
	"github.com/machester4/jane/bktree"
	"github.com/machester4/jane/chain"
	"github.com/machester4/jane/constants"
)


func addRecommendsBK(w *Word, tree bktree.BKTree) func()  {
	return func() {
		results := tree.Search(chain.BkWord(w.Value), constants.MaxDistance)
		for _, result := range results {
			w.Recommends = append(w.Recommends, result.Entry.(chain.BkWord))
		}
	}
}

// Contexts are dictionaries with the respective words.
// have higher priority than the words of the language itself
func Recommend(chain *Chain, lang string, tree bktree.BKTree) {
	words := chain.Words
	var steps []func()

	for _, word := range words {
		steps = append(steps, addRecommendsBK(word, tree))
	}

	stages := []*Stage{
		{
			Name:  "Add words recommends",
			Steps: steps,
		},
		{
			Name:  "Add articles recommends",
			Steps: nil,
		},
		{
			Name:  "Add punct recommends",
			Steps: nil,
		},
	}
	pipeline := Pipeline{Stages: stages}
	pipeline.Start(false)
}
