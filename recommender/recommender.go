package recommender

import (
	"fmt"

	"github.com/agnivade/levenshtein"
)

func addRecommends(word *Word) func() {
	return func() {
		distance := levenshtein.ComputeDistance(word.Value, "hol")
		fmt.Printf("distance %d - %s\n", distance, word.Value)
		word.Recommends = [3]string{word.Value, string(distance)}
	}
}

// Contexts are dictionaries with the respective words.
// have higher priority than the words of the language itself
func Recommend(chain *Chain, lang string, contexts []string) {
	// wordsDict := helpers.GetDictionary("chivito")
	words := chain.Words
	var steps []func()

	for _, word := range words {
		steps = append(steps, addRecommends(word))
	}

	stages := []*Stage{
		{
			Name:  "Add words recommends",
			Steps: steps,
		},
		{
			Name:  "Add articles recommends",
			Steps: steps,
		},
		{
			Name:  "Add punct recommends",
			Steps: steps,
		},
	}
	pipeline := Pipeline{Stages: stages}
	pipeline.Start(false)
}
