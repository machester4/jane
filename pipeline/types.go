package pipeline

import (
	"sync"

	"github.com/machester4/jane/chain"
)

type Word = chain.Word

type Stage struct {
	Name   string
	Worker func(w *[]Word)
}

type WordPipeline struct {
	wg     sync.WaitGroup
	Words  *[]Word
	Stages *[]Stage
}
