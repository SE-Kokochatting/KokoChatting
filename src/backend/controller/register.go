package controller

import (
	"KokoChatting/global"
	"KokoChatting/model/req"
	"KokoChatting/model/res"
	"KokoChatting/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RegisterController struct {
	baseController
	registerService *service.RegisterService
}

// Register [POST]
// PATH: api/v1/user/register
// Function: 实现用户的注册
func (regCtl *RegisterController) Register(c *gin.Context) {
	Regreq := &req.UserReq{}
	if err := c.BindJSON(Regreq); err != nil {
		global.Logger.Error("bind json error", zap.Error(err))
	}
	// return corresponding error
	uid, err := regCtl.registerService.Register(Regreq.Name, Regreq.Password)
	if err != nil {
		regCtl.WithErr(global.RegisterError, c)
	}

	Regres := &res.UserRes{}
	Regres.Data.Uid = uid
	regCtl.WithData(Regres, c)
}

func NewRegisterController() *RegisterController {
	return &RegisterController{
		registerService: service.NewRegisterService(),
	}
}