/**
* @Author: Jason
* @Date: 2020/9/30 16:18
* @File : game_account
* @Software: GoLand
**/

package main

import (
	"context"
	"dygame/RpcModule"
	"dygame/common"
	"dygame/dcore"
	"dygame/gameserver/manager"
	protoMsg "dygame/message"

	"github.com/golang/protobuf/proto"
)

func accountLogin(ses *dcore.WsSession, message proto.Message) {
	l := message.(*protoMsg.C2S_Login)

	// 调用DB登录接口
	plr, err := RpcModule.DbRpcClient.PlayerDBLogin(context.Background(), &protoMsg.L2D_Login{
		Account:  l.GetAccount(),
		Password: l.GetPassword(),
	})
	if err != nil {
		common.OperateTip(ses, err.Error(), 0)
		return
	}
	loginOk(ses, plr)
}

func loginOk(ses *dcore.WsSession, plr *protoMsg.D2L_UserInfo) {
	// 表示已经存在帐号
	if plr != nil {
		player := manager.PlayerMgr.GetPlayerByID(plr.GameID)
		if player == nil {
			player = &common.PlayerItem{
				Ses: ses,
			}
		}
	}
}
