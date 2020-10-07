/**
* @Author: Jason
* @Date: 2020/9/29 16:13
* @File : db
* @Software: GoLand
**/

package dcore

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

type SqlCon struct {
	DbEngine *xorm.Engine
}

func (self *SqlCon) Connect(constr string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", constr)
	self.DbEngine = engine
	if err != nil {
		log.WithFields(log.Fields{
			"db": "mysql",
		}).Error(err)
		return nil
	}

	engine.SetConnMaxLifetime(time.Second * 500)
	engine.SetMaxOpenConns(100)
	engine.SetMapper(names.SameMapper{})

	if err := self.DbEngine.Ping(); err != nil {
		fmt.Println(err)
		return nil
	}

	return engine
}
