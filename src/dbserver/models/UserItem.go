package models

type UserItem struct {
	ID        int   `xorm:"not null pk autoincr INT(11)"`
	UID       int64 `xorm:"not null BIGINT(64)"`
	GameID    int64 `xorm:"not null BIGINT(64)"`
	ItemID    int   `xorm:"not null INT(11)"`
	ItemType  int   `xorm:"not null INT(11)"`
	ItemNum   int   `xorm:"not null default 0 INT(11)"`
	PlaceType int   `xorm:"default 0 comment('存放类型,仓库还是背包') INT(11)"`
}
