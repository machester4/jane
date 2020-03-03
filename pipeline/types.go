package pipeline

import "github.com/machester4/jane/chain"

// Alias for bind chain types to local types
type Chain = chain.Chain
type Block = chain.Block

// NOTE: delayed pipes run after no-delayed pipes

type BlockPipe struct {
	Name    string
	Delayed bool
	Task    func(block *Block)
}

type ChainPipe struct {
	Name    string
	Delayed bool
	Task    func(chain *Chain)
}

type Pipeline struct {
	chain             *Chain
	chainPipes        []ChainPipe
	blockPipes        []BlockPipe
	chainPipesDelayed []*ChainPipe
	blockPipesDelayed []*BlockPipe
}
