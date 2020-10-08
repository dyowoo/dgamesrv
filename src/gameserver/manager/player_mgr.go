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

func (m *playerManager) AddItem(player *common.PlayerItem) {
	m.arrayPlayer[player.GameID] = player
}

func (m *playerManager) IsPlayerExist(gameID int64) bool {
	_, ok := m.arrayPlayer[gameID]
	return ok
}

func (m *playerManager) Delete(player *common.PlayerItem) {
	delete(m.arrayPlayer, player.GameID)
}

func (m *playerManager) GetPlayerByID(gameID int64) (player *common.PlayerItem) {
	v, ok := m.arrayPlayer[gameID]
	if !ok {
		return nil
	}
	return v
}

func (m *playerManager) GetPlayerBySesID(sesID int64) (player *common.PlayerItem) {
	for _, value := range m.arrayPlayer {
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
