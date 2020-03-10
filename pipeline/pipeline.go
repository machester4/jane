package pipeline

import "fmt"

func (p *WordPipeline) logDone(stgName string) {
	fmt.Printf("Stage %q done all steps!\n", stgName)
}

func (s *Stage) logDone(sIndex int) {
	fmt.Printf("[%q]: Step -> %d Done!\n", s.Name, sIndex)
}

func (s *Stage) run(word *Word) {
	for i, step := range s.Steps {
		step(word)
		s.logDone(i)
	}
	// Done all Steps of Stage, stage done
}

func (p *WordPipeline) runnerSync() {
	for _, word := range p.Words {
		for _, stage := range p.Stages {
			stage.run(word)
			p.logDone(stage.Name)
		}
	}
}

func (p *WordPipeline) runnerAsync()  {
	p.wg.Add(len(p.Words) * len(p.Stages))

	for _, word := range p.Words {
		for _, stage := range p.Stages {
			go (func() {
				defer p.wg.Done()
				stage.run(word)
				p.logDone(stage.Name)
			})()
		}
	}
	p.wg.Wait()
}

func (p *WordPipeline) Start(sync bool) {
	if sync {
		p.runnerSync()
	} else {
		p.runnerAsync()
	}
}
