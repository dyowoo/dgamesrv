package models

type Account struct {
	ID       int    `xorm:"not null pk autoincr INT(32)"`
	GameID   int64  `xorm:"not null BIGINT(64)"`
	Account  string `xorm:"not null VARCHAR(32)"`
	Name     string `xorm:"not null default '''' VARCHAR(32)"`
	Password string `xorm:"not null VARCHAR(16)"`
}
