package provider

import (
	"KokoChatting/global"
	"KokoChatting/model/dataobject"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"  // mysql driver (can not delete)
	"go.uber.org/zap"
)

type storeUserProfile struct {
	Uid uint64 `gorm:"column:uid;primary_key;AUTO_INCREMENT"`
	Name string `gorm:"column:name"`
	Password string `gorm:"column:password"`
}

type mysqlProvider struct{
	mysqlDb *gorm.DB
}

// 获取mysql中存储的用户信息表
func (mysql *mysqlProvider) getStoreUserProfile(userprofile *dataobject.UserProfile) (*storeUserProfile, error) {
	user := &storeUserProfile{}
	user.Uid = userprofile.Uid
	user.Name = userprofile.Name
	user.Password = userprofile.Password
	return user, nil
}

func NewMysqlProvider() *mysqlProvider {
	db, err := MysqlInit()
	if err != nil{
		global.Logger.Error("init mysql error", zap.Error(err))
		panic(err)
	}
	return &mysqlProvider{
		mysqlDb: db,
	}
}