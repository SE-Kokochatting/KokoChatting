package service

import (
	"KokoChatting/global"
	"KokoChatting/model/dataobject"
	"KokoChatting/provider"
	"KokoChatting/util"
	"fmt"
	"go.uber.org/zap"
)

type UserService struct {
	userProvider *provider.UserProvider
}

// Register 用户注册处理逻辑
func (userSrv *UserService) Register(name string, password string) (uint64, error) {
	encryptedPassword := util.Encryption(password)
	// 将该用户添加到数据库表中，然后返回该用户表中的uid
	// 同时进行检查该用户是否可以添加到数据库中（查重）
	userProfile := &dataobject.UserProfile{
		Name: name,
		Password: encryptedPassword,
	}
	if err := userSrv.userProvider.CheckExist(userProfile); err != nil {
		global.Logger.Error("add user error", zap.Error(err))
		return 0, err
	}

	uid, err := userSrv.userProvider.CreateUser(userProfile)
	if err != nil {
		global.Logger.Error("add user error", zap.Error(err))
		return 0, err
	}

	return uid, err
}

// Login 用户登录处理逻辑：检查该用户uid是否存在，如果存在，检查对应密码是否正确，返回布尔结果和err
func (userSrv *UserService) Login(uid uint64, password string) (bool, error) {
	userProfile := &dataobject.UserProfile{
		Uid: uid,
	}
	// check the user is existed and find the info
	if err := userSrv.userProvider.CheckExist(userProfile); err != nil {
		global.Logger.Error(fmt.Sprintf("user is not existed, uid: %d", uid), zap.Error(err))
		return false, err
	}
	if password != userProfile.Password {
		global.Logger.Error(fmt.Sprintf("user password is invalid, uid: %d", uid), zap.Error(global.PasswordError))
		return false, global.PasswordError
	}
	return true, nil
}

func NewRegisterService() *UserService {
	return &UserService{
		userProvider: provider.NewRegisterProvider(),
	}
}