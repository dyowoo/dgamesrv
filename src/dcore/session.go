package dcore

import (
	"fmt"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
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
func (s *WsSession) SetID(id int64) {
	s.id = id
}

/**
获取会话ID
*/
func (s *WsSession) GetID() int64 {
	return s.id
}

/**
关闭会话
*/
func (s *WsSession) Close() {
	var msg Event
	msg.EType = EVENT_CLOSE
	msg.Ses = s
	s.sendPipe.Add(msg)
}

/**
发送封包
*/
func (s *WsSession) Send(protocolType uint32, buffer proto.Message) {
	var msg Event
	data, err := proto.Marshal(buffer)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	msg.EType = EVENT_SEND
	msg.EData = EncodeMsg(protocolType, data)
	s.sendPipe.Add(msg)
}

/**
发送循环
*/
func (s *WsSession) sendLoop() {
	for s.conn != nil {
		msg := s.sendPipe.Get()
		if msg.EType == EVENT_CLOSE {
			break
		}

		err := s.conn.WriteMessage(websocket.BinaryMessage, msg.EData)
		if err != nil {
			log.WithFields(log.Fields{
				"WsSession": "SendLoop",
			}).Error(err)
			break
		}
	}

	// 关闭连接
	if s.conn != nil {
		_ = s.conn.Close()
		s.conn = nil
	}

	// 通知完成
	s.exitSync.Done()
}

/**
接收循环
*/
func (s *WsSession) recvLoop() {
	for s.conn != nil {
		_, raw, err := s.conn.ReadMessage()
		e := Event{
			EVENT_RECV,
			raw,
			s,
		}

		if err != nil {
			if !IsEOFOrNetReadError(err) {
				log.WithFields(log.Fields{
					"WsSession": "recvLoop",
				}).Error(err)
			}

			//s.Close()

			var msg Event
			msg.EType = EVENT_CLOSE
			msg.Ses = s
			s.pInstance.RecvPostEvent(msg)

			break
		}

		s.pInstance.RecvPostEvent(e)
	}

	s.exitSync.Done()
}

/**
启动会话
*/
func (s *WsSession) Start() {
	// 将会话添加到会话管理器
	s.pInstance.sesMgr.Add(s)

	// 需要接收和发送协程同时完成时才算真正完成
	s.exitSync.Add(2)

	go func() {
		// 等待2个协程任务结束
		s.exitSync.Wait()

		// 将自己的会话从管理器中移除
		s.pInstance.sesMgr.Remove(s)
	}()

	// 启动接收goroutine
	go s.recvLoop()

	// 启动发送goroutine
	go s.sendLoop()
}

/**
创建一个会话
*/
func NewSession(conn *websocket.Conn, p *NetComponent) *WsSession {
	var sendPipe pipe
	sendPipe.chSend = make(chan Event, 100)

	session := &WsSession{
		conn:      conn,
		pInstance: p,
		sendPipe:  &sendPipe,
	}

	return session
}
