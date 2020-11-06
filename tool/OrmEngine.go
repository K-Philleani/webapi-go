package tool

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"webapi-go/model")

var DbEngine *Orm

type Orm struct {
	*xorm.Engine
}
const URL = "root:123456@tcp(124.70.71.78:3306)/webApi"

func OrmEngine(cfg *Config) (*Orm, error){
	engine, err := xorm.NewEngine(cfg.DataBase.Driver, URL)
	if err != nil {
		return nil, err
	}
	engine.ShowSQL(true)
	err = engine.Sync2(new(model.SmsCode))
	if err != nil {
		return nil, err
	}
	orm := new(Orm)
	orm.Engine = engine
	DbEngine = orm
	return orm, nil
}
