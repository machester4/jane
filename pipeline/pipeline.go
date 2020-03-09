package pipeline

import "fmt"

func getDelayedPipesChain(pipes []PipeChain) []*PipeChain {
	result := make([]*PipeChain, 0)
	for _, p := range pipes {
		if p.Delayed {
			result = append(result, &p)
		}
	}
	return result
}

func getDelayedPipesBlock(pipes []PipeBlock) []*PipeBlock {
	result := make([]*PipeBlock, 0)
	for _, p := range pipes {
		if p.Delayed {
			result = append(result, &p)
		}
	}
	return result
}

func getDelayedPipesWord(pipes []PipeWord) []*PipeWord {
	result := make([]*PipeWord, 0)
	for _, p := range pipes {
		if p.Delayed {
			result = append(result, &p)
		}
	}
	return result
}

func (p *Pipeline) AddChainPipe(cp PipeChain) {
	p.chainPipes = append(p.chainPipes, cp)
	p.chainPipesDelayed = getDelayedPipesChain(p.chainPipes)
}

func (p *Pipeline) AddBlockPipe(bp PipeBlock) {
	p.blockPipes = append(p.blockPipes, bp)
	p.blockPipesDelayed = getDelayedPipesBlock(p.blockPipes)
}

func (p *Pipeline) AddWordPipe(wp PipeWord) {
	p.wordPipes = append(p.wordPipes, wp)
	p.wordPipesDelayed = getDelayedPipesWord(p.wordPipes)
}

func (p *Pipeline) Run() (words []Word) {
	words = p.chain.GetWords()
	p.wg.Add(len(words))

	for _, word := range words {
		for _, wp := range p.wordPipes {
			go wp.Task(&word, &p.wg)
		}
	}
	p.wg.Wait()
	fmt.Printf("Words en run %q\n", words)
	return
}

func New(chain *Chain) Pipeline {
	return Pipeline{
		chain: chain,
	}
}
