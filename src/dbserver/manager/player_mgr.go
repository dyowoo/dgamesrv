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

func (m *DBPlayerMgr) AddPlayer(player *common.PlayerItem) {
	m.arrPlayer[player.GameID] = player
}

func (m *DBPlayerMgr) IsPlayerExist(gameID int64) bool {
	_, ok := m.arrPlayer[gameID]
	return ok
}

func (m *DBPlayerMgr) Delete(player *common.PlayerItem) {
	delete(m.arrPlayer, player.GameID)
}

func (m *DBPlayerMgr) GetPlayer(gameID int64) (player *common.PlayerItem) {
	v, ok := m.arrPlayer[gameID]
	if !ok {
		return nil
	}
	return v
}

func (m *DBPlayerMgr) GetPlayerByAccount(account string) (player *common.PlayerItem) {
	for _, v := range m.arrPlayer {
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
