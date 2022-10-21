package controller

import (
	"KokoChatting/global"

	"github.com/gin-gonic/gin"
)

type baseController struct{
	
}

func (b *baseController) WithErr(ApiErr error, c *gin.Context) {
	err := ApiErr.(global.Error)
	c.JSON(200, gin.H{
		"code": err.Status,
		"msg": err.Error(),
		"data": nil,
	})
}

func (b *baseController) WithData(data interface{}, c *gin.Context) {
	c.JSON(200, data)
}




