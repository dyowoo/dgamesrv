package models

type UserHero struct {
	ID             int    `xorm:"not null pk autoincr INT(11)"`
	UID            int64  `xorm:"not null BIGINT(64)"`
	HeroID         int    `xorm:"not null comment('英雄ID') INT(11)"`
	GameID         int64  `xorm:"not null comment('所属玩家') BIGINT(64)"`
	Lv             int    `xorm:"not null comment('等级') INT(11)"`
	Exp            int    `xorm:"not null comment('经验') INT(11)"`
	Hp             int    `xorm:"not null default 100 comment('血量') INT(11)"`
	Att            int    `xorm:"not null comment('攻击力') INT(11)"`
	Mana           int    `xorm:"not null default 0 comment('魔法值') INT(11)"`
	AttackDistance int    `xorm:"not null default 50 INT(11)"`
	Skills         string `xorm:"not null default '' comment('技能数组') VARCHAR(255)"`
}
