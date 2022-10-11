package provider

import (
	"KokoChatting/global"
	"go.uber.org/zap"
)

type ExampleProvider struct{
	mysqlProvider
}

func (prd *ExampleProvider) ExampleCRUD(args ...interface{}) error {
	prd.mysqlDb.Table("example").Where("dummy = ?",1) // 表为空，会报错，仅为示例
	if prd.mysqlDb.Error != nil{
		// log
		global.Logger.Error("some error msg",zap.Error(prd.mysqlDb.Error))
	}
	return global.ConfigPathError
}

func NewExampleProvider()*ExampleProvider{
	return &ExampleProvider{
		mysqlProvider{},
	}
}