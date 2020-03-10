package pipeline

import (
	"fmt"
)

func (s *Stage) logDone(stgName string) {
	fmt.Printf("Stage %q done all steps!\n", stgName)
}

func (s *Stage) logStepDone(sIndex int) {
	fmt.Printf("[%q]: Step -> %d Done!\n", s.Name, sIndex)
}

func (p *Pipeline) runnerSync() {
	for _, stage := range p.Stages {
		for i, step := range stage.Steps {
			step()
			stage.logStepDone(i)
		}
		stage.logDone(stage.Name)
	}
}

func (p *Pipeline) runnerAsync()  {
	for _, stage := range p.Stages {
		// Wait current stage run all step in parallel
		p.wg.Add(len(stage.Steps))
		for i, step := range stage.Steps {
			go (func() {
				step()
				stage.logStepDone(i)
				p.wg.Done()
			})()
		}
		p.wg.Wait()
		stage.logDone(stage.Name)
	}
}

// Run all stages sync mode
func (p *Pipeline) Start(sync bool) {
	if sync {
		p.runnerSync()
	} else {
		p.runnerAsync()
	}
}
