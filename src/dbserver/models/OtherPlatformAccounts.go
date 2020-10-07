package models

type OtherPlatformAccounts struct {
	UnionID    string `xorm:"not null pk VARCHAR(32)"`
	OpenID     string `xorm:"not null VARCHAR(32)"`
	GameID     int64  `xorm:"not null BIGINT(64)"`
	PlatformID int    `xorm:"not null default 0 INT(11)"`
}
