package models

type EquipmentAttrEx struct {
	ID      int   `xorm:"not null pk autoincr INT(11)"`
	UID     int64 `xorm:"not null pk comment('装备UID') BIGINT(64)"`
	GameID  int64 `xorm:"not null pk BIGINT(64)"`
	TxID    int   `xorm:"not null pk default 0 comment('特殊属性ID') INT(11)"`
	TxAttr1 int   `xorm:"not null default 0 comment('特殊属性值1') INT(11)"`
	TxAttr2 int   `xorm:"not null default 0 comment('特殊属性值2') INT(11)"`
}
