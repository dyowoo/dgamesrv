package dcore

import (
	"sync"
	"sync/atomic"
)

type SessionManager struct {
	sesMap   sync.Map //使用ID关联会话
	sesIDGen int64    //记录已经生成的会话ID流水号
	count    int32    //记录当前在使用的会话数量
}

func (self *SessionManager) SetIDBase(base int64) {
	atomic.StoreInt64(&self.sesIDGen, base)
}

func (self *SessionManager) Count() int32 {
	return self.count
}

func (self *SessionManager) Add(ses *WsSession) {
	id := atomic.AddInt64(&self.sesIDGen, 1)
	atomic.AddInt32(&self.count, 1)
	ses.SetID(id)
	self.sesMap.Store(id, ses)
}

func (self *SessionManager) Remvoe(ses *WsSession) {
	self.sesMap.Delete(ses.GetID())
	atomic.AddInt32(&self.count, -1)
}

func (self *SessionManager) GetSession(id int64) *WsSession {
	if val, ok := self.sesMap.Load(id); ok {
		return val.(*WsSession)
	}
	return nil
}

func (self *SessionManager) RangeSession(callback func(session *WsSession) bool) {
	self.sesMap.Range(func(key, value interface{}) bool {
		return callback(value.(*WsSession))
	})
}

func (self *SessionManager) CloseAllSession() {
	self.RangeSession(func(session *WsSession) bool {
		session.Close()
		return true
	})
}

/**
获取活跃的会话数量
*/
func (self *SessionManager) SessionCount() int32 {
	count := atomic.LoadInt32(&self.count)
	return count
}
