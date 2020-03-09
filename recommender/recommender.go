package recommender

import (
	"fmt"
	"sync"

	"github.com/machester4/jane/pipeline"
)

// Contexts are dictionaries with the respective words.
// have higher priority than the words of the language itself
func Recommend(chain *Chain, lang string, contexts []string) (words []Word) {
	mainPipeline := pipeline.New(chain)

	wordPipe := pipeline.PipeWord{
		Delayed: false,
		Task: func(w *pipeline.Word, wg *sync.WaitGroup) {
			defer wg.Done()
			recs := [3]string{"a", "b", "c"}
			w.Recommendations = recs
			fmt.Printf("dentro del task %q\n", w.Recommendations)
		},
	}

	mainPipeline.AddWordPipe(wordPipe)
	words = mainPipeline.Run()
	return words
}
