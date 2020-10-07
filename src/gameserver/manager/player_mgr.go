/**
* @Author: Jason
* @Date: 2020/9/30 16:44
* @File : player_mgr
* @Software: GoLand
**/

package manager

import "dygame/common"

type ArrayPlayer map[int64]*common.PlayerItem
type playerManager struct {
	arrayPlayer ArrayPlayer
}

func (self *playerManager) AddItem(player *common.PlayerItem) {
	self.arrayPlayer[player.GameID] = player
}

func (self *playerManager) IsPlayerExist(gameID int64) bool {
	_, ok := self.arrayPlayer[gameID]
	return ok
}

func (self *playerManager) Delete(player *common.PlayerItem) {
	delete(self.arrayPlayer, player.GameID)
}

func (self *playerManager) GetPlayerByID(gameID int64) (player *common.PlayerItem) {
	v, ok := self.arrayPlayer[gameID]
	if !ok {
		return nil
	}
	return v
}

func (self *playerManager) GetPlayerBySesID(sesID int64) (player *common.PlayerItem) {
	for _, value := range self.arrayPlayer {
		if value.Ses.GetID() == sesID {
			return value
		}
	}

	return nil
}

var PlayerMgr playerManager

func init() {
	PlayerMgr = playerManager{
		arrayPlayer: make(ArrayPlayer, 0),
	}
}
