package pipeline

func getDelayedCP(pipes []ChainPipe) []*ChainPipe {
	result := make([]*ChainPipe, 0)
	for _, p := range pipes {
		if p.Delayed {
			result = append(result, &p)
		}
	}
	return result
}

func getDelayedBP(pipes []BlockPipe) []*BlockPipe {
	result := make([]*BlockPipe, 0)
	for _, p := range pipes {
		if p.Delayed {
			result = append(result, &p)
		}
	}
	return result
}

func (p *Pipeline) AddChainPipe(cp ChainPipe) {
	p.chainPipes = append(p.chainPipes, cp)
	p.chainPipesDelayed = getDelayedCP(p.chainPipes)
}

func (p *Pipeline) AddBlockPipe(bp BlockPipe) {
	p.blockPipes = append(p.blockPipes, bp)
	p.blockPipesDelayed = getDelayedBP(p.blockPipes)
}

func New(chain *Chain) Pipeline {
	return Pipeline{
		chain: chain,
	}
}
