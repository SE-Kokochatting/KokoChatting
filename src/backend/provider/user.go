package provider

import (
	"KokoChatting/global"
	"KokoChatting/model/dataobject"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type RegisterProvider struct {
	mysqlProvider
}

// AddUser 增加用户到用户表中
func (regPro *RegisterProvider) AddUser(name string, encrytedpassword string) (uint64, error) {
	// 查找有无重复的name用户
	var userProfileEntity = &dataobject.UserProfile{
		Name: name,
		Password: encrytedpassword,
	}

	// check whether the name is useful
	dbClient := regPro.mysqlProvider.mysqlDb
	err := dbClient.Where("name = ?", name).First(userProfileEntity).Error
	if err != gorm.ErrRecordNotFound {
		global.Logger.Error("name is already used", zap.Error(err))
		return 0, err
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		global.Logger.Error("register error", zap.Error(err))
		return 0, err
	}

	if err := dbClient.Create(userProfileEntity).Error; err != nil {
		global.Logger.Error("add user error", zap.Error(err))
		return 0, err
	}

	// find the inserted user and return the uid
	err = dbClient.Where("name = ?", name).First(userProfileEntity).Error
	if err != nil {
		global.Logger.Error("register error when find the inserted user", zap.Error(err))
		return 0, err
	}
	// dbClient.Debug().Create(userProfileEntity)
	return userProfileEntity.Uid, nil
}

func NewRegisterProvider() *RegisterProvider {
	return &RegisterProvider{
		mysqlProvider: *NewMysqlProvider(),
	}
}