package dcore

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"sync"
)

/**
会话
*/
type WsSession struct {
	conn      *websocket.Conn // websocket连接
	pInstance *NetComponent   // 网络实例
	sendPipe  *pipe           // 发送管道
	exitSync  sync.WaitGroup  // 退出同步器
	id        int64           // 会话ID
}

/**
设置会话ID
*/
func (self *WsSession) SetID(id int64) {
	self.id = id
}

/**
获取会话ID
*/
func (self *WsSession) GetID() int64 {
	return self.id
}

/**
关闭会话
*/
func (self *WsSession) Close() {
	var msg Event
	msg.EType = EVENT_CLOSE
	msg.Ses = self
	self.sendPipe.Add(msg)
}

/**
发送封包
*/
func (self *WsSession) Send(protocolType uint32, buffer proto.Message) {
	var msg Event
	data, err := proto.Marshal(buffer)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	msg.EType = EVENT_SEND
	msg.EData = EncodeMsg(protocolType, data)
	self.sendPipe.Add(msg)
}

/**
发送循环
*/
func (self *WsSession) sendLoop() {
	for self.conn != nil {
		msg := self.sendPipe.Get()
		if msg.EType == EVENT_CLOSE {
			break
		}

		err := self.conn.WriteMessage(websocket.BinaryMessage, msg.EData)
		if err != nil {
			log.WithFields(log.Fields{
				"WsSession": "SendLoop",
			}).Error(err)
			break
		}
	}

	// 关闭连接
	if self.conn != nil {
		_ = self.conn.Close()
		self.conn = nil
	}

	// 通知完成
	self.exitSync.Done()
}

/**
接收循环
*/
func (self *WsSession) recvLoop() {
	for self.conn != nil {
		_, raw, err := self.conn.ReadMessage()
		e := Event{
			EVENT_RECV,
			raw,
			self,
		}

		if err != nil {
			if !IsEOFOrNetReadError(err) {
				log.WithFields(log.Fields{
					"WsSession": "recvLoop",
				}).Error(err)
			}

			//self.Close()

			var msg Event
			msg.EType = EVENT_CLOSE
			msg.Ses = self
			self.pInstance.RecvPostEvent(msg)

			break
		}

		self.pInstance.RecvPostEvent(e)
	}

	self.exitSync.Done()
}

/**
启动会话
*/
func (self *WsSession) Start() {
	// 将会话添加到会话管理器
	self.pInstance.sesMgr.Add(self)

	// 需要接收和发送协程同时完成时才算真正完成
	self.exitSync.Add(2)

	go func() {
		// 等待2个协程任务结束
		self.exitSync.Wait()

		// 将自己的会话从管理器中移除
		self.pInstance.sesMgr.Remvoe(self)
	}()

	// 启动接收goroutine
	go self.recvLoop()

	// 启动发送goroutine
	go self.sendLoop()
}

/**
创建一个会话
*/
func NewSession(conn *websocket.Conn, p *NetComponent) *WsSession {
	var sendPipe pipe
	sendPipe.ch_send = make(chan Event, 100)

	session := &WsSession{
		conn:      conn,
		pInstance: p,
		sendPipe:  &sendPipe,
	}

	return session
}
