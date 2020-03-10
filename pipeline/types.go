package pipeline

import (
	"sync"

	"github.com/machester4/jane/chain"
)

type Word = chain.Word

type Stage struct {
	Name   string
	Steps []func(w *Word)
	Lifo bool
}

type WordPipeline struct {
	wg     sync.WaitGroup
	Words  []*Word
	Stages []*Stage
}