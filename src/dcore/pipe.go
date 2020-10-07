package dcore

type pipe struct {
	ch_send chan Event
	cb      PipeCb
}

func (self *pipe) Add(e Event) {
	self.ch_send <- e
}

func (self *pipe) Get() Event {
	return <-self.ch_send
}

func (self *pipe) Start() {
	go func() {
		for {
			var data Event
			select {
			case data = <-self.ch_send:
				self.cb(data)
			}
		}
	}()
}
