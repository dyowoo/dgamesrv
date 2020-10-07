/**
* @Author: Jason
* @Date: 2020/9/30 16:22
* @File : system
* @Software: GoLand
**/

package common

import (
	"dygame/dcore"
	protoMsg "dygame/message"
)

func OperateTip(ses *dcore.WsSession, err string, code int32) {
	info := protoMsg.S2C_OperateTip{
		Msg:     err,
		ErrCode: code,
	}
	ses.Send(uint32(protoMsg.CMD_LOGIN_Resp_OperateTip), &info)
}
