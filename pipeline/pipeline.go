package pipeline

import (
	"fmt"
	"sync"
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

func (s *Stage) wrapRoutine(wg *sync.WaitGroup, index int) {
	defer wg.Done()
	s.Steps[index]()
	s.logStepDone(index)
}

func (p *Pipeline) runnerAsync() {
	for _, stage := range p.Stages {
		// Wait current stage run all step in parallel
		p.wg.Add(len(stage.Steps))
		for i, _ := range stage.Steps {
			go stage.wrapRoutine(&p.wg, i)
		}
		p.wg.Wait()
		stage.logDone(stage.Name)
	}
}

// Stages are executed synchronously,
// and steps are executed depending on the value of the sync variable
func (p *Pipeline) Start(sync bool) {
	if sync {
		p.runnerSync()
	} else {
		p.runnerAsync()
	}
}
