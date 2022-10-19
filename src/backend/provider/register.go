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
	var userProfileEntity = &dataobject.StoreUserProfile{}
	dbClient := regPro.mysqlProvider.mysqlDb
	if err := dbClient.Where("name = ?", name).First(userProfileEntity).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			global.Logger.Error("name is already used", zap.Error(err))
			return 0, err
		}
	}

	// 添加到数据表中
	userProfileEntity.Name = name
	userProfileEntity.Password = encrytedpassword
	if err := dbClient.Create(userProfileEntity).Error; err != nil {
		global.Logger.Error("add user error", zap.Error(err))
		return 0, err
	}
	return userProfileEntity.Uid+100000, nil
}

func NewRegisterProvider() *RegisterProvider {
	return &RegisterProvider{
		mysqlProvider: *NewMysqlProvider(),
	}
}