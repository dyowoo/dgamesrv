package dcore

type EVENT_TYPE int32
type PipeCb func(Event)

const (
	EVENT_NULL  EVENT_TYPE = iota //未定义
	EVENT_SEND                    //socket发送事件
	EVENT_RECV                    //socket接收事件
	EVENT_CLOSE                   //socket关闭
)
