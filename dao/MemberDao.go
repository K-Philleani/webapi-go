package dao

import (
	"log"
	"webapi-go/model"
	"webapi-go/tool"
)

type MemberDao struct {
	*tool.Orm
}


func (md *MemberDao) InsertCode(sms model.SmsCode) int64{
	res, err := md.Engine.InsertOne(&sms)
	if err != nil {
		log.Println(err)
	}
	return res
}