/**
* @Author: Jason
* @Date: 2020/9/30 10:06
* @File : playerMgr
* @Software: GoLand
**/

package manager

import "dygame/common"

type ArrayPlayer map[int64]*common.PlayerItem
type DBPlayerMgr struct {
	arrPlayer ArrayPlayer
}

func (self *DBPlayerMgr) AddPlayer(player *common.PlayerItem) {
	self.arrPlayer[player.GameID] = player
}

func (self *DBPlayerMgr) IsPlayerExist(gameID int64) bool {
	_, ok := self.arrPlayer[gameID]
	return ok
}

func (self *DBPlayerMgr) Delete(player *common.PlayerItem) {
	delete(self.arrPlayer, player.GameID)
}

func (self *DBPlayerMgr) GetPlayer(gameID int64) (player *common.PlayerItem) {
	v, ok := self.arrPlayer[gameID]
	if !ok {
		return nil
	}
	return v
}

func (self *DBPlayerMgr) GetPlayerByAccount(account string) (player *common.PlayerItem) {
	for _, v := range self.arrPlayer {
		if v.Account == account {
			return v
		}
	}
	return nil
}

var PlayerMgr DBPlayerMgr

func init() {
	PlayerMgr = DBPlayerMgr{
		arrPlayer: make(ArrayPlayer, 0),
	}
}
