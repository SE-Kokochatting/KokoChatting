package provider

import (
	"KokoChatting/global"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver (can not delete)
	"go.uber.org/zap"
)

type mysqlProvider struct{
	mysqlDb *gorm.DB
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