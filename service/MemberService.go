package service

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"log"
	"math/rand"
	"time"
	"webapi-go/dao"
	"webapi-go/model"
	"webapi-go/tool"
)

type MemberService struct {

}

func (ms *MemberService) SendCode(phone string) bool{
	// 1.产生一个验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	// 2.调用阿里云sdk 完成发送
	config := tool.GetConfig().Sms
	client, err := dysmsapi.NewClientWithAccessKey(config.RegionId, config.AccessKey, config.AccessSecret)
	if err != nil {
		log.Println(err)
		return false
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phone
	request.SignName = config.SignName
	request.TemplateCode = config.TemplateCode
	par, err := json.Marshal(map[string]interface{}{
		"code": code,
	})
	request.TemplateParam = string(par)

	// 3.接收返回结果，并判断发送状态
	response, err := client.SendSms(request)
	fmt.Print(response)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if response.Code == "OK" {
		// 将验证码保存到数据库中
		smsCode := model.SmsCode{
			Phone:      phone,
			BizId:      response.BizId,
			Code:       code,
			CreateTime: time.Now().Unix(),
		}
		memberDao := dao.MemberDao{ tool.DbEngine }
		res := memberDao.InsertCode(smsCode)
		log.Println("发送的验证码:", code)
		return res > 0
	}
	return false
}


