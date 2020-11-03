package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type DbConfig struct {
	DbName  string 	 `json:"db_name"`
	DbUser  string   `json:"db_user"`
	DbPwd   string   `json:"db_pwd"`
	DbTable string   `json:"db_table"`
	Address string	 `json:"address"`
}

var _DbCfg *DbConfig

func ParseDbConfig(path string) (*DbConfig, error){
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&_DbCfg)
	if err != nil {
		return nil, err
	}
	return _DbCfg, nil
}
