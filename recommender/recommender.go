package recommender

import (
	"fmt"
	"github.com/agnivade/levenshtein"
)

func addRecommends(word *Word)  {
	fmt.Println(word.Value)
	levenshtein.ComputeDistance(word.ToString(), "word en dic")
	word.Recommendations = [3]string{word.ToString(),"c"}
}

// Contexts are dictionaries with the respective words.
// have higher priority than the words of the language itself
func Recommend(chain *Chain, lang string, contexts []string) []*Word {
	// wordsDict := helpers.GetDictionary("chivito")
	words := chain.GetWords()

	stages := []*Stage{
		{
			Name: "Recommend",
			Steps: []func(w *Word){addRecommends},
		},
	}
	pipeline := WordPipeline{Words: words, Stages: stages}
	pipeline.Start(false)

	return words
}
