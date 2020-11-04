package main

import (
	"github.com/gin-gonic/gin"
	"webapi-go/controller"
	"webapi-go/tool"
)

func main() {
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	registerRouter(router)
	router.Run(cfg.AppHost + ":" + cfg.AppPort)
}

// 路由设置
func registerRouter(router *gin.Engine) {
	new(controller.EnterController).Router(router)
	new(controller.MemberController).Router(router)
}
