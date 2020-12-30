/**
* @Author: Jason
* @Date: 2020/9/29 16:13
* @File : db
* @Software: GoLand
**/

package dcore

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

type SqlCon struct {
	DbEngine *xorm.Engine
}

func (c *SqlCon) Connect(constr string) *xorm.Engine {
	// 创建数据库引擎
	engine, err := xorm.NewEngine("mysql", constr)
	c.DbEngine = engine
	if err != nil {
		log.WithFields(log.Fields{
			"db": "mysql",
		}).Error(err)
		return nil
	}

	// 配置数据库参数
	engine.SetConnMaxLifetime(time.Second * 500)
	engine.SetMaxOpenConns(100)
	engine.SetMapper(names.SameMapper{})

	if err := c.DbEngine.Ping(); err != nil {
		fmt.Println(err)
		return nil
	}

	return engine
}
