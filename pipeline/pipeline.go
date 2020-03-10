package pipeline

import "fmt"

func (p *WordPipeline) logAndDone(stgName string) {
	fmt.Println("Done Stage ->", stgName)
	p.wg.Done()
}

func (p *WordPipeline) runner(sync bool) {
	if sync {
		for _, stage := range *p.Stages {
			(func() {
				stage.Worker(p.Words)
				defer p.logAndDone(stage.Name)
			})()
		}
	} else {
		for _, stage := range *p.Stages {
			go (func() {
				stage.Worker(p.Words)
				defer p.logAndDone(stage.Name)
			})()
		}
	}
}

func (p *WordPipeline) Start() {
	p.wg.Add(len(*p.Stages))
	p.runner(false)
	p.wg.Wait()
}
