package models

type PlayerAttribute struct {
	ID        int   `xorm:"not null pk autoincr INT(11)"`
	GameID    int64 `xorm:"not null default 0 BIGINT(64)"`
	Lv        int   `xorm:"not null comment('等级') INT(11)"`
	Exp       int   `xorm:"not null comment('经验') INT(11)"`
	OutHeroID int64 `xorm:"not null default 0 BIGINT(64)"`
	MapUID    int64 `xorm:"not null BIGINT(64)"`
}
