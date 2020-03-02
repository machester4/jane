package pipeline

import "github.com/machester4/jane/chain"

// NOTE: delayed pipes run after no-delayed pipes

type Pipe interface {
	isDelayed() bool
}

type BlockPipe struct {
	name    string
	delayed bool
	task    func(block *chain.Block)
}

type ChainPipe struct {
	name    string
	delayed bool
	task    func(block *chain.Chain)
}

type Pepeline struct {
	chain *chain.Chain
	chainPipes []ChainPipe
	blockPipes []BlockPipe
	chainPipesDelayed []*ChainPipe
	blockPipesDelayed []*BlockPipe
}