package provider

import (
	"KokoChatting/model/dataobject"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"  // mysql driver (can not delete)
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
	db,err := gorm.Open("mysql","user:password@/dbname?charset=utf8&parseTime=True&loc=Local")  // todo :mysql 选项信息从配置文件中解析
	if err != nil{
		// log and panic
	}
	return &mysqlProvider{
		mysqlDb: db,
	}
}