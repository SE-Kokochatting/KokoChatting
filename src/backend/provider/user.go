package provider

import (
	"KokoChatting/global"
	"KokoChatting/model/dataobject"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type UserProvider struct {
	mysqlProvider
}

// CheckExist 检查用户是否存在，存在直接补全对应的结构体信息，不存在返回nil
func (UserPro *UserProvider) CheckExist(userprofile *dataobject.UserProfile) error {
	// if uid is zero, search by name; if uid is not zero, search by uid
	dbClient := UserPro.mysqlProvider.mysqlDb
	if 0 == userprofile.Uid {
		err := dbClient.Where("name = ?", userprofile.Name).First(userprofile).Error
		if err != gorm.ErrRecordNotFound {
			global.Logger.Error("name is already used", zap.Error(err))
			return err
		}
		if err != nil && err != gorm.ErrRecordNotFound {
			global.Logger.Error("register error", zap.Error(err))
			return err
		}
	} else {
		err := dbClient.Where("uid = ?", userprofile.Uid).First(userprofile).Error
		if err != nil{
			global.Logger.Error("login error: user is not exist", zap.Error(err))
			return err
		}
	}
	return nil
}

func (UserPro *UserProvider) CreateUser(userprofile *dataobject.UserProfile) (uint64, error){
	dbClient := UserPro.mysqlProvider.mysqlDb
	if err := dbClient.Create(userprofile).Error; err != nil {
		global.Logger.Error("add user error", zap.Error(err))
		return 0, err
	}

	// find the inserted user and return the uid
	err := dbClient.Where("name = ?", userprofile.Name).First(userprofile).Error
	if err != nil {
		global.Logger.Error("register error when find the inserted user", zap.Error(err))
		return 0, err
	}
	// dbClient.Debug().Create(userProfileEntity)
	return userprofile.Uid, nil
}

// UpdateEntry update all columns by uid
func (UserPro *UserProvider) UpdateEntry(lastedProfile *dataobject.UserProfile) error {
	dbClient := UserPro.mysqlProvider.mysqlDb

	userprofile := &dataobject.UserProfile{}
	err := dbClient.Where("uid = ?", lastedProfile.Uid).First(userprofile).Error
	if err != nil{
		global.Logger.Error("update entry error: user is not exist", zap.Error(err))
		return err
	}

	// update the entry
	err = dbClient.Model(&dataobject.UserProfile{}).Where("uid=?", lastedProfile.Uid).Updates(dataobject.UserProfile{
		Password: lastedProfile.Password,
		Name: lastedProfile.Name,
		AvatarUrl: lastedProfile.AvatarUrl,
	}).Error
	if err != nil {
		global.Logger.Error("update entry error: can not update", zap.Error(err))
		return err
	}

	return nil
}

func NewUserProvider() *UserProvider {
	return &UserProvider{
		mysqlProvider: *NewMysqlProvider(),
	}
}