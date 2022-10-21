package service

import (
	"KokoChatting/global"
	"KokoChatting/provider"
	"KokoChatting/util"
	"go.uber.org/zap"
)

type RegisterService struct {
	registerProver *provider.RegisterProvider
}

// Register 用户注册接口
func (regSrv *RegisterService) Register(name string, password string) (uint64, error) {
	encryptedPassword := util.Encryption(password)
	// 将该用户添加到数据库表中，然后返回该用户表中的uid
	// 同时进行检查该用户是否可以添加到数据库中（查重）
	uid, err := regSrv.registerProver.AddUser(name, encryptedPassword)
	if err != nil {
		global.Logger.Error("add user error", zap.Error(err))
		return 0, err
	}

	return uid, err
}

func NewRegisterService() *RegisterService {
	return &RegisterService{
		registerProver: provider.NewRegisterProvider(),
	}
}