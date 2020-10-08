package dcore

import (
	"net/http"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
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

func (c *NetComponent) Start(cb PipeCb) error {
	c.corePipe = &pipe{
		ch_send: make(chan Event, 100),
		cb:      cb,
	}
	c.corePipe.Start()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		conn, err := upGrader.Upgrade(writer, request, nil)
		if err != nil {
			log.WithFields(log.Fields{
				"net": "HandleFunc",
			}).Error("Upgrade fail:", err)
			return
		}

		ses := NewSession(conn, c)
		ses.Start()
	})

	err := http.ListenAndServe(c.Constr, nil)
	if err != nil {
		log.WithFields(log.Fields{
			"net": "ListenAndServe",
		}).Error(err)
		return err
	}

	return nil
}

func (c *NetComponent) RecvPostEvent(e Event) {
	c.corePipe.Add(e)
}
