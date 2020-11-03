package controller

import (
	"github.com/gin-gonic/gin"
	"webapi-go/model"
)

func GetAccountInfo(c *gin.Context) {
	var account model.Account
	rowList, err := account.GetAccountInfo()
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"code": 1,
		"dataList": rowList,
	})
}
