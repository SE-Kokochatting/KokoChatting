package controller

import (
	"KokoChatting/global"
	"KokoChatting/model/res"
	"KokoChatting/service"
	"github.com/gin-gonic/gin"
)

type RegisterController struct {
	baseController
	registerService *service.RegisterService
}

// Register [POST]
// PATH: api/v1/user/register
// Function: 实现用户的注册
func (regCtl *RegisterController) Register(c *gin.Context, req *res.UserRes) {
	req = &res.UserRes{}

	name := c.Param("name")
	password := c.Param("password")
	// return corresponding error
	uid, err := regCtl.registerService.Register(name, password)
	if err != nil {
		regCtl.WithErr(global.Error{
			Status: 1000,
			Err: err,
		}, c)
	}

	req.Data.Uid = uid
	regCtl.WithData(req, c)
}

func NewRegisterController() *RegisterController {
	return &RegisterController{
		registerService: service.NewRegisterService(),
	}
}