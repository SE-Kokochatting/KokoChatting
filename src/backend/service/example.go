package service

import "KokoChatting/provider"


type ExampleService struct{
	exampleProvider *provider.ExampleProvider
}

func (srv *ExampleService) Example(args ...interface{}) error {
	srv.exampleProvider.ExampleCRUD(args)
	return nil
}


func NewExampleService()*ExampleService{
	return &ExampleService{
		exampleProvider: provider.NewExampleProvider(),
	}
}
