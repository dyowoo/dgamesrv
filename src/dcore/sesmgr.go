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

func (m *SessionManager) SetIDBase(base int64) {
	atomic.StoreInt64(&m.sesIDGen, base)
}

func (m *SessionManager) Count() int32 {
	return m.count
}

func (m *SessionManager) Add(ses *WsSession) {
	id := atomic.AddInt64(&m.sesIDGen, 1)
	atomic.AddInt32(&m.count, 1)
	ses.SetID(id)
	m.sesMap.Store(id, ses)
}

func (m *SessionManager) Remvoe(ses *WsSession) {
	m.sesMap.Delete(ses.GetID())
	atomic.AddInt32(&m.count, -1)
}

func (m *SessionManager) GetSession(id int64) *WsSession {
	if val, ok := m.sesMap.Load(id); ok {
		return val.(*WsSession)
	}
	return nil
}

func (m *SessionManager) RangeSession(callback func(session *WsSession) bool) {
	m.sesMap.Range(func(key, value interface{}) bool {
		return callback(value.(*WsSession))
	})
}

func (m *SessionManager) CloseAllSession() {
	m.RangeSession(func(session *WsSession) bool {
		session.Close()
		return true
	})
}

/**
获取活跃的会话数量
*/
func (m *SessionManager) SessionCount() int32 {
	count := atomic.LoadInt32(&m.count)
	return count
}
