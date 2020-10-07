/**
* @Author: Jason
* @Date: 2020/9/30 9:56
* @File : server_account
* @Software: GoLand
**/

package main

import (
	"context"
	"dygame/common"
	"dygame/dbserver/manager"
	protoMsg "dygame/message"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func (s *server) PlayerDBLogin(_ context.Context, in *protoMsg.L2D_Login) (*protoMsg.D2L_UserInfo, error) {
	account := in.GetAccount()
	fmt.Println("playerLogin :" + account)
	player := manager.PlayerMgr.GetPlayerByAccount(account)
	var err error
	if player == nil {
		player, err = getAccount(in.Account)
		if err != nil {
			log.WithFields(log.Fields{
				"module": "dbserver",
			}).Info(" get account by db fail:", err)
			return &protoMsg.D2L_UserInfo{}, err
		}
	}
	return loginComplete(player)
}

func loginComplete(player *common.PlayerItem) (*protoMsg.D2L_UserInfo, error) {
	d2lUserInfo := &protoMsg.D2L_UserInfo{
		GameID:   player.GameID,
		NickName: player.NickName,
		Score:    player.Gold,
		Password: player.Password,
		Account:  player.Account,
		Lv:       player.Lv,
	}

	return d2lUserInfo, nil
}
