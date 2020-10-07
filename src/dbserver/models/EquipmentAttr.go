package models

type EquipmentAttr struct {
	ID     int   `xorm:"not null pk autoincr INT(11)"`
	UID    int64 `xorm:"not null BIGINT(64)"`
	GameID int64 `xorm:"not null BIGINT(64)"`
	Lv     int   `xorm:"not null default 0 comment('等级') INT(11)"`
	Exp    int   `xorm:"not null default 0 comment('经验') INT(11)"`
	Att    int   `xorm:"not null default 0 comment('攻击力') INT(11)"`
	Def    int   `xorm:"not null default 0 comment('防御值') INT(11)"`
	Speed  int   `xorm:"not null default 0 comment('速度') INT(11)"`
	Mz     int   `xorm:"not null default 0 comment('命中') INT(11)"`
	Hp     int   `xorm:"not null default 0 comment('血量值') INT(11)"`
	Mp     int   `xorm:"not null default 0 comment('魔法值') INT(11)"`
}
