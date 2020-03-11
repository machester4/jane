package recommender

import (
	"github.com/machester4/jane/constants"

	"github.com/machester4/jane/helpers"

	"github.com/agnivade/levenshtein"
)

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

// Contexts are dictionaries with the respective words.
// have higher priority than the words of the language itself
func Recommend(chain *Chain, lang string, contexts []string) {
	wordsDict := helpers.GetDictionary(lang)
	words := chain.Words
	var steps []func()

	for _, word := range words {
		steps = append(steps, addRecommends(word, wordsDict))
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
