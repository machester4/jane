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

func New(chain *Chain, chainPipes []ChainPipe, blockPipes []BlockPipe) (pipeline Pipeline) {
	pipeline = Pipeline{
		chain:             chain,
		chainPipes:        chainPipes,
		blockPipes:        blockPipes,
		chainPipesDelayed: getDelayedCP(chainPipes),
		blockPipesDelayed: getDelayedBP(blockPipes),
	}
	return
}
