package provider

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"  // mysql driver (can not delete)
)

type mysqlProvider struct{
	mysqlDb *gorm.DB
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