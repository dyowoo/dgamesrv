package dcore

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type NetComponent struct {
	Constr   string //连接字符串
	corePipe *pipe
	sesMgr   SessionManager
}

var upGrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (self *NetComponent) Start(cb PipeCb) error {
	self.corePipe = &pipe{
		ch_send: make(chan Event, 100),
		cb:      cb,
	}
	self.corePipe.Start()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		conn, err := upGrader.Upgrade(writer, request, nil)
		if err != nil {
			log.WithFields(log.Fields{
				"net": "HandleFunc",
			}).Error("Upgrade fail:", err)
			return
		}

		ses := NewSession(conn, self)
		ses.Start()
	})

	err := http.ListenAndServe(self.Constr, nil)
	if err != nil {
		log.WithFields(log.Fields{
			"net": "ListenAndServe",
		}).Error(err)
		return err
	}

	return nil
}

func (self *NetComponent) RecvPostEvent(e Event) {
	self.corePipe.Add(e)
}
