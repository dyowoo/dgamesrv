/**
* @Author: Jason
* @Date: 2020/9/30 10:10
* @File : player_item
* @Software: GoLand
**/

package common

import (
	"dygame/dcore"
	"math"
)

func (p Position) Distance(target Position) int32 {
	dx := math.Abs(float64(p.X) - float64(target.X))
	dy := math.Abs(float64(p.Y) - float64(target.Y))
	return int32(math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2)))
}

type PlayerItem struct {
	Ses      *dcore.WsSession
	Account  string
	NickName string
	Password string
	Gold     int64
	GameID   int64
	Lv       int32
	Exp      int64
	MapID    int32
}
