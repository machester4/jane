package pipeline

import "github.com/machester4/jane/chain"

// delayed pipes run after no-delayed pipes

type BlockPipe struct {
	name    string
	delayed bool
	task    func(block *chain.Block)
}

type ChainBlock struct {
	name    string
	delayed bool
	task    func(block *chain.Chain)
}
