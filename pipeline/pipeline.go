package pipeline

import "github.com/machester4/jane/chain"

func getDelayed(pipes Pipe)  {
	
}

func (p *Pepeline) New(chain *chain.Chain, chainPipes []ChainPipe, blockPipes []BlockPipe) (pepeline Pepeline) {
	pepeline = Pepeline{
		chain: chain,
		chainPipes:chainPipes,
		blockPipes:blockPipes,
	}
	
	return
}