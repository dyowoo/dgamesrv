package models

type GameScore struct {
	GameID int64 `xorm:"not null pk BIGINT(64)"`
	Gold   int64 `xorm:"not null default 0 BIGINT(64)"`
}
