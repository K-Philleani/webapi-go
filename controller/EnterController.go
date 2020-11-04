package controller

import "github.com/gin-gonic/gin"

type EnterController struct {

}

func (enter *EnterController) Router(engine *gin.Engine) {
	engine.GET("/enter",enter.Enter )
}

func (enter *EnterController) Enter(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"message": "this is a enter",
	})
}

