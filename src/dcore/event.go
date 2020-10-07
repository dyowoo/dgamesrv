package dcore

type Event struct {
	EType EVENT_TYPE
	EData []byte
	Ses   *WsSession
}
