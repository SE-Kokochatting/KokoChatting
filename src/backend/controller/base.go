package controller

import (
	"KokoChatting/global"

	"github.com/gin-gonic/gin"
)

type baseController struct{
	
}

func (b *baseController) WithErr(ApiErr error, c *gin.Context) {
	err := ApiErr.(global.Error)
	c.JSON(err.Status, gin.H{
		"Err": err.Error(),
	})
}

func (b *baseController) WithData(data interface{}, c *gin.Context) {
	c.JSON(200, data)
}




