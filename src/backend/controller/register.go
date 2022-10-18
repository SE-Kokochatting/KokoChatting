package controller

import (
	"KokoChatting/global"
	"KokoChatting/model/dataobject"
	"KokoChatting/model/req"
	"KokoChatting/model/res"
	"KokoChatting/model/utilstruct"
	"KokoChatting/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"time"
)

type RegisterController struct {
	baseController
	registerService *service.RegisterService
}

// Register [POST]
// PATH: api/v1/user/register
// Function: 实现用户的注册
func (regCtl *RegisterController) Register(c *gin.Context) {
	regreq := &req.UserReq{}
	if err := c.BindJSON(regreq); err != nil {
		global.Logger.Error("bind json error", zap.Error(err))
	}
	// return corresponding error
	uid, err := regCtl.registerService.Register(regreq.Name, regreq.Password)
	if err != nil {
		regCtl.WithErr(global.RegisterError, c)
	}

	regres := &res.UserRes{}
	regres.Data.Uid = uid

	// 生成token信息
	claims := utilstruct.Claims{
		UserProfile: dataobject.UserProfile{
			Uid: uid,
			Password: regreq.Password,
			Name: regreq.Name,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3*time.Hour*time.Duration(1))),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	config := global.Config{}
	mySecret := config.GetConfigByName("jwt.secret").(string)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySecret)
	if err != nil {
		global.Logger.Error("generate jwt error", zap.Error(err))
		regCtl.WithErr(global.RegisterError, c)
	}
	regres.Data.Token = tokenString
	regCtl.WithData(regres, c)
}

func NewRegisterController() *RegisterController {
	return &RegisterController{
		registerService: service.NewRegisterService(),
	}
}