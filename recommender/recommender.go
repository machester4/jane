package recommender

import (
	"fmt"
	"github.com/machester4/jane/bktree"
	"github.com/machester4/jane/constants"

	"github.com/agnivade/levenshtein"
)

type BkWord string
func (x BkWord) Distance(e bktree.Entry) int {
	a := string(x)
	b := string(e.(BkWord))

	return levenshtein.ComputeDistance(a, b)
}

func addRecommends(word *Word, dicWords []string) func() {
	return func() {
		for _, wd := range dicWords {
			distance := levenshtein.ComputeDistance(word.Value, wd)
			if wd == word.Value {
				word.Recommends = []string{}
				break
			}
			if distance < constants.MaxDistance {
				// fmt.Printf("distance %d - %s\n", distance, wd)
				word.Recommends = append(word.Recommends, wd)
			}
		}
	}
}

func addRecommendsBK(w *Word, tree bktree.BKTree) func()  {
	return func() {
		results := tree.Search(BkWord(w.Value), constants.MaxDistance)
		for _, result := range results {
			// w.Recommends = append(w.Recommends, string(result.Entry))
			fmt.Printf("\t%s (distance: %d)\n", result.Entry.(BkWord), result.Distance)
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
