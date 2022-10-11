package service

import (
	"KokoChatting/global"
	"KokoChatting/provider"
	"go.uber.org/zap"
)


type ExampleService struct{
	exampleProvider *provider.ExampleProvider
}

func (srv *ExampleService) Example(args ...interface{}) error {
	err := srv.exampleProvider.ExampleCRUD(args)
	if err != nil{
		global.Logger.Error("some error msg",zap.Error(err))
		return err
	}
	return nil
}

func NewExampleService()*ExampleService{
	return &ExampleService{
		exampleProvider: provider.NewExampleProvider(),
	}
}
