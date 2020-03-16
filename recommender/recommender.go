package recommender

import (
	"fmt"
	"github.com/machester4/jane/bktree"
	"github.com/machester4/jane/constants"
	"github.com/machester4/jane/helpers"

	"github.com/agnivade/levenshtein"
)

type word string
func (x word) Distance(e bktree.Entry) int {
	a := string(x)
	b := string(e.(word))

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
		results := tree.Search(word(w.Value), constants.MaxDistance)
		for _, result := range results {
			// w.Recommends = append(w.Recommends, string(result.Entry))
			fmt.Printf("\t%s (distance: %d)\n", result.Entry.(word), result.Distance)
		}
	}
}

// Contexts are dictionaries with the respective words.
// have higher priority than the words of the language itself
func Recommend(chain *Chain, lang string, contexts []string) {
	words := chain.Words
	wordsDict := helpers.GetDictionary(lang)
	var steps []func()
	var tree bktree.BKTree
	for _, w := range wordsDict {
		tree.Add(word(w))
	}

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
