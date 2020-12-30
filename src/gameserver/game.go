/**
* @Author: Jason
* @Date: 2020/9/29 15:32
* @File : gameServer
* @Software: GoLand
**/

package main

import (
	"dygame/RpcModule"
	"dygame/dcore"
	protoMsg "dygame/message"
	"fmt"
	"math/rand"
	"reflect"
	"runtime/debug"
	"time"

	"github.com/golang/protobuf/proto"
)

type msgCb func(ses *dcore.WsSession, message proto.Message)

type MessageProc struct {
	funcMap     map[uint32]msgCb
	funcNameMap map[uint32]string
}

/**
注册消息回调
*/
func (p MessageProc) RegisterCallback(msgType uint32, cb msgCb, messageName string) {
	_, ok := p.funcMap[msgType]
	if ok {
		fmt.Printf("msgType:%d, msgName:%s 已存在\n", msgType, messageName)
		return
	}
	p.funcMap[msgType] = cb
	p.funcNameMap[msgType] = messageName
}

/**
消息回调处理
*/
func (p MessageProc) RunCallback(ses *dcore.WsSession, msgType uint32, msgByte []byte) {
	_, ok := p.funcMap[msgType]
	if !ok {
		fmt.Printf("msgType:%d 不存在\n", msgType)
		return
	}

	msgRef := proto.MessageType(p.funcNameMap[msgType])
	msg := reflect.New(msgRef.Elem()).Interface().(proto.Message)
	_ = proto.Unmarshal(msgByte, msg)

	p.funcMap[msgType](ses, msg)
}

var cbMap MessageProc

var net = dcore.NetComponent{
	Constr: "0.0.0.0:1920",
}

func main() {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 初始化RPC
	RpcModule.InitRpc()

	cbMap = MessageProc{
		funcMap:     make(map[uint32]msgCb),
		funcNameMap: make(map[uint32]string),
	}

	// 注册回调
	cbMap.RegisterCallback(uint32(protoMsg.CMD_LOGIN_Req_AccountLogin), accountLogin, "protoMsg.C2S_Login") // 帐号登录

	_ = net.Start(msgHandle)
}

func msgHandle(e dcore.Event) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
			fmt.Println(string(debug.Stack()))
		}
	}()

	switch e.EType {
	case dcore.EVENT_RECV:
		msgType, _, msgData := dcore.DecodeMsg(e.EData)
		cbMap.RunCallback(e.Ses, msgType, msgData)
	case dcore.EVENT_CLOSE:

	}
}
