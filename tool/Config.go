package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName string  	`json:"app_name"`
	AppMode string 		`json:"app_mode"`
	AppHost string 		`json:"app_host"`
	AppPort string  	`json:"app_port"`
	Sms 	SmsConfig 	`json:"sms"`
}

type SmsConfig struct {
	Signname 	  string `json:"sign_name"`
	TemplateCode  string `json:"template_code"`
	RegionId 	  string `json:"region_id"`
	AppKey 		  string `json:"app_key"`
	AppSecret 	  string `json:"app_secret"`
}

var _cfg *Config = nil

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&_cfg);
	if err!= nil {
		return nil, err
	}
	return _cfg, nil
}