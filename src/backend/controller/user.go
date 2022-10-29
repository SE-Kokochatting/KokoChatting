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

type UserController struct {
	baseController
	userService *service.UserService
}

// Register [POST]
// PATH: api/v1/user/register
// Function: 实现用户的注册
func (userCtl *UserController) Register(c *gin.Context) {
	regReq := &req.UserRegisterReq{}
	if err := c.BindJSON(regReq); err != nil {
		global.Logger.Error("bind json error", zap.Error(err))
	}

	// return corresponding error
	uid, err := userCtl.userService.Register(regReq.Name, regReq.Password)
	if err != nil {
		userCtl.WithErr(global.RegisterError, c)
		return
	}

	regRes := &res.UserRegisterRes{}
	regRes.Data.Uid = uid

	userCtl.WithData(regRes, c)
}

// Login [POST]
// PATH: api/v1/user/login
// Function: 实现用户的登录
func (userCtl *UserController) Login(c *gin.Context) {
	loginReq := &req.UserLoginReq{}
	if err := c.BindJSON(loginReq); err != nil {
		global.Logger.Error("login bind json error", zap.Error(err))
	}

	// 用户登录逻辑：
	// controller层：首先拿到用户uid和对应的密码
	// service层：向上提供验证函数，检查该用户uid是否存在，如果存在，检查对应密码是否正确，返回布尔结果和err
	// provider层：向上提供以下函数：根据profile表检查用户是否存在，若存在，返回对应的用户信息，不存在，返回nil

	isProper, err := userCtl.userService.Login(loginReq.Uid, loginReq.Password)
	if false == isProper {
		global.Logger.Error("login err", zap.Error(err))
		userCtl.WithErr(global.LoginError, c)
		return
	}

	loginRes := &res.UserLoginRes{}
	// generate tokenStr
	claims := utilstruct.Claims{
		Uid: loginReq.Uid,
		Password: loginReq.Password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3*time.Hour*time.Duration(1))),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	config := global.GetGlobalConfig()
	mySecret := config.GetConfigByName("jwt.secret").(string)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(mySecret))
	if err != nil {
		global.Logger.Error("generate jwt error", zap.Error(err))
		userCtl.WithErr(global.LoginError, c)
		return
	}
	loginRes.Data.Token = tokenString

	userCtl.WithData(loginRes, c)
}

// GetUserInfo [GET]
// PATH: api/v1/user
// Function: get detailed information of user
func (userCtl *UserController) GetUserInfo(c *gin.Context) {
	userInfoReq := &req.UserInfoReq{}
	if err := c.BindJSON(userInfoReq); err != nil {
		global.Logger.Error("get user info bind json error", zap.Error(err))
	}

	userProfile := &dataobject.UserProfile{
		Uid: userInfoReq.Uid,
	}

	err := userCtl.userService.GetUserInfo(userInfoReq.Uid, userProfile)
	if err != nil {
		global.Logger.Error("get user info error", zap.Error(err))
		userCtl.WithErr(global.GetInfoError, c)
		return
	}

	userInfoRes := &res.UserInfoRes{
		Data: struct {
			Uid uint64 `json:"uid"`
			Name string `json:"name"`
			AvatarUrl string `json:"avatarUrl"`
		}{
			Uid: userProfile.Uid,
			Name: userProfile.Name,
			AvatarUrl: userProfile.AvatarUrl,
		},
	}

	userCtl.WithData(userInfoRes, c)
}

// SetUserAvatar [POST]
// PATH: api/v1/user/avatar
// Function: set avatarUrl for the user
func (userCtl *UserController) SetUserAvatar(c *gin.Context) {
	uid := userCtl.getUid(c)

	setAvatarReq := &req.UserSetAvatarReq{}
	if err := c.BindJSON(setAvatarReq); err != nil {
		global.Logger.Error("set avatar bind json error", zap.Error(err))
	}

	err := userCtl.userService.SetAvatar(uid, setAvatarReq.AvatarUrl)
	if err != nil {
		global.Logger.Error("set avatar error", zap.Error(err))
		userCtl.WithErr(global.AvatarError, c)
		return
	}

	setAvatarRes := &res.UserSetAvatarRes{}

	userCtl.WithData(setAvatarRes, c)
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}