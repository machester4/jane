package lib

import (
	"fmt"
	"sync"
)

type stage struct {
	name  string
	steps []func()
	lifo  bool // TODO: Implement
}

type pipeline struct {
	wg     sync.WaitGroup
	stages []*stage
}

func (s *stage) logDone(stgName string) {
	fmt.Printf("Stage %q done all steps!\n", stgName)
}

func (s *stage) logStepDone(sIndex int) {
	fmt.Printf("[%q]: Step -> %d Done!\n", s.name, sIndex)
}

func (p *pipeline) runnerSync() {
	for _, stage := range p.stages {
		for i, step := range stage.steps {
			step()
			stage.logStepDone(i)
		}
		stage.logDone(stage.name)
	}
}

func (s *stage) wrapRoutine(wg *sync.WaitGroup, index int) {
	defer wg.Done()
	s.steps[index]()
	s.logStepDone(index)
}

func (p *pipeline) runnerAsync() {
	for _, stage := range p.stages {
		// Wait current stage run all step in parallel
		p.wg.Add(len(stage.steps))
		for i := range stage.steps {
			go stage.wrapRoutine(&p.wg, i)
		}
		p.wg.Wait()
		stage.logDone(stage.name)
	}
}

// Stages are executed synchronously,
// and steps are executed depending on the value of the sync variable
func (p *pipeline) start(sync bool) {
	if sync {
		p.runnerSync()
	} else {
		p.runnerAsync()
	}
}
