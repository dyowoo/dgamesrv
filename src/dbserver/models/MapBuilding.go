package models

import (
	"time"
)

type MapBuilding struct {
	ID           int       `xorm:"not null pk autoincr INT(11)"`
	MapUID       int64     `xorm:"not null pk BIGINT(64)"`
	GameID       int64     `xorm:"not null BIGINT(64)"`
	BuildingID   int       `xorm:"not null default 0 comment('建筑ID') INT(11)"`
	PositionX    int       `xorm:"not null INT(11)"`
	PositionY    int       `xorm:"not null INT(1)"`
	BuildingUID  int64     `xorm:"not null default 0 BIGINT(64)"`
	BuildingTime time.Time `xorm:"not null default CURRENT_TIMESTAMP DATETIME"`
	IsDelete     int       `xorm:"not null default 0 TINYINT(4)"`
}
