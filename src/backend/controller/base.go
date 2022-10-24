package controller

import (
	"KokoChatting/global"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type baseController struct{
	
}

func (b *baseController) WithErr(ApiErr error, c *gin.Context) {
	err := ApiErr.(global.Error)
	c.JSON(404, gin.H{
		"code": err.Status,
		"msg": err.Error(),
		"data": nil,
	})
}

func (b *baseController) WithData(data interface{}, c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg": "successful",
		"data": data,
	})
}
// getUid 帮助有鉴权要求的接口获得鉴权
func (b *baseController) getUid (c *gin.Context) uint64 {
	value,exist := c.Get("userUid")
	if exist{
		uid,ok := value.(float64)
		if ok{
			return uint64(uid)
		}
		global.Logger.Error("value conv to uint64 error")
		return 0
	}
	global.Logger.Error("there is no uid",zap.Error(errors.New("there is no uid")))
	return 0
}




