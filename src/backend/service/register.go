package service

import (
	"KokoChatting/global"
	"KokoChatting/provider"
	"KokoChatting/util"
)

type RegisterService struct {
	registerProver *provider.RegisterProvider
}

// Register 用户注册接口
func (regSrv *RegisterService) Register(name string, password string) (uint64, error) {
	encryptedPassword := util.Encryption(password)

}

func NewRegisterService() *RegisterService {
	return &RegisterService{
		registerProver: provider.NewRegisterProvider(),
	}
}