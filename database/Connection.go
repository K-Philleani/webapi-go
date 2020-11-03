package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"webapi-go/tool"
)

var DB *sql.DB


func init() {
	var err error
	dbCfg, err := tool.ParseDbConfig("./config/database.json")
	if err != nil {
		panic(err)
	}
	DSN := dbCfg.DbUser+ ":" + dbCfg.DbPwd + "@tcp(" + dbCfg.Address + ")/" + dbCfg.DbTable
	DB, err = sql.Open(dbCfg.DbName, DSN)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("数据库连接成功")
}
