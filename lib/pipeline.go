package lib

type chmsg struct {
	stginx int
	data   *field
}

type stage struct {
	name   string
	worker func(*field)
}

type pipeline struct {
	ch     chan chmsg
	tlimit int
	fields []*field
	stages []*stage
}

func (p *pipeline) runner(stginx int, f *field) {
	p.stages[stginx].worker(f)
	p.ch <- chmsg{stginx: stginx, data: f}
}

func (p *pipeline) observer() {
	for msg := range p.ch {
		p.tlimit--
		if msg.stginx != (len(p.stages) - 1) {
			go p.runner(msg.stginx+1, msg.data)
		}
		if p.tlimit == 0 {
			close(p.ch)
		}
	}
}

func (p *pipeline) run() {
	p.ch = make(chan chmsg)
	p.tlimit = len(p.fields) * len(p.stages)

	for _, f := range p.fields {
		go p.runner(0, f)
	}
	p.observer()
}
