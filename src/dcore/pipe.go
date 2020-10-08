package dcore

type pipe struct {
	ch_send chan Event
	cb      PipeCb
}

func (p *pipe) Add(e Event) {
	p.ch_send <- e
}

func (p *pipe) Get() Event {
	return <-p.ch_send
}

func (p *pipe) Start() {
	go func() {
		for {
			var data Event
			select {
			case data = <-p.ch_send:
				p.cb(data)
			}
		}
	}()
}
