package provider

import (
	"KokoChatting/global"
	"go.uber.org/zap"
)

type ExampleProvider struct{
	mysqlProvider
}

func (prd *ExampleProvider) ExampleCRUD(args ...interface{}) error {
	prd.mysqlDb.Table("example").Where("dummy = ?",1)
	if prd.mysqlDb.Error != nil{
		// log
		global.Logger.Error("some error msg",zap.Error(prd.mysqlDb.Error))
	}
	return prd.mysqlDb.Error
}

func NewExampleProvider()*ExampleProvider{
	return &ExampleProvider{
		mysqlProvider{},
	}
}