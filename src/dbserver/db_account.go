/**
* @Author: Jason
* @Date: 2020/9/30 11:38
* @File : db_account
* @Software: GoLand
**/

package main

import (
	"dygame/common"
	"dygame/dbserver/manager"
	"dygame/dbserver/models"
	log "github.com/sirupsen/logrus"
)

func getAccount(account string) (*common.PlayerItem, error) {
	user := models.Account{
		Account: account,
	}
	has, err := con.DbEngine.Get(&user)
	if err != nil {
		return &common.PlayerItem{}, err
	}
	if has {
		player := manager.PlayerMgr.GetPlayer(user.GameID)
		if player == nil {
			player = &common.PlayerItem{
				Account:  user.Account,
				NickName: user.Name,
				Password: user.Password,
				Gold:     0,
				GameID:   user.GameID,
				Lv:       0,
				Exp:      0,
				MapID:    0,
			}
			gameScore := models.GameScore{
				GameID: player.GameID,
			}
			has, err = con.DbEngine.Get(&gameScore)
			if err != nil {
				log.WithFields(log.Fields{
					"Models": "GameScore",
				}).Error("get score by db fail:", err)
			}
			if has {
				player.Gold = gameScore.Gold
			}
			manager.PlayerMgr.AddPlayer(player)
		}
		return player, nil
	}
	return nil, nil
}
