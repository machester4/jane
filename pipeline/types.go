package pipeline

import (
	"sync"

	"github.com/machester4/jane/chain"
)

// Alias for bind chain types to local types
type Chain = chain.Chain
type Block = chain.Block
type Word = chain.Word

// NOTE: delayed pipes run after no-delayed pipes

type PipeChain struct {
	Delayed bool
	Task    func(c Chain)
}

type PipeBlock struct {
	Delayed bool
	Task    func(b Block)
}

type PipeWord struct {
	Delayed bool
	Task    func(w *Word, wg *sync.WaitGroup)
}

type Pipeline struct {
	wg                sync.WaitGroup
	chain             *Chain
	chainPipes        []PipeChain
	blockPipes        []PipeBlock
	wordPipes         []PipeWord
	chainPipesDelayed []*PipeChain
	blockPipesDelayed []*PipeBlock
	wordPipesDelayed  []*PipeWord
}
