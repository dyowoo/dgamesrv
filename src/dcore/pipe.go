package dcore

/**
管道结构
*/
type pipe struct {
	chSend chan Event
	cb     PipeCb
}

func (p *pipe) Add(e Event) {
	p.chSend <- e
}

func (p *pipe) Get() Event {
	return <-p.chSend
}

func (p *pipe) Start() {
	go func() {
		for {
			var data Event
			select {
			case data = <-p.chSend:
				p.cb(data)
			}
		}
	}()
}
