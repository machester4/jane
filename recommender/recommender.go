package recommender

import "fmt"

// Contexts are dictionaries with the respective words.
// have higher priority than the words of the language itself
func Recommend(chain *Chain, lang string, contexts []string) []Word {
	words := chain.GetWords()
	recommendStage := Stage{Name: "Recommend", Worker: func(w *[]Word) {
		for _, word := range words {
			fmt.Println(word.Value)
		}
	}}
	stages := []Stage{recommendStage}
	pipeline := WordPipeline{Words: &words, Stages: &stages}
	pipeline.Start()

	return words
}
