package recommender

import "fmt"

// Contexts are dictionaries with the respective words.
// have higher priority than the words of the language itself
func New(chain *Chain, lang string, contexts []string) (words []Word) {
	words = chain.GetWords()

	for _, w := range words {
		fmt.Printf("add recommends of %q\n", w.Value)
	}
	return words
}
