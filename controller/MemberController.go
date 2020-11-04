package controller

import (
	"github.com/gin-gonic/gin"
	"webapi-go/service"
)

type MemberController struct {

}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendCode", mc.sendSmsCode)
}


func (mc *MemberController) sendSmsCode(c *gin.Context) {
	// 发送验证码
	phone, exist := c.GetQuery("phone")
	if !exist {
		c.JSON(200, map[string]interface{} {
			"code": 0,
			"message": "参数解析失败",
		})
		return
	}
	service := service.MemberService{}
	isSend := service.SendCode(phone)
	if isSend {
		c.JSON(200, map[string]interface{}{
			"code": 1,
			"message": "发送成功",
		})
	} else {
		c.JSON(200, map[string]interface{}{
			"code": 0,
			"message": "发送失败",
		})
	}
}



