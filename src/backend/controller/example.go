package controller

import (
	"KokoChatting/service"

	"github.com/gin-gonic/gin"
)

type ExampleController struct{
	baseController
	exampleService *service.ExampleService
}

func (controller *ExampleController) Example(c *gin.Context){
	err := controller.exampleService.Example(1,2,3,"test","example")
	if err != nil{
		controller.WithErr(err, c)
		return
	}
	// controller.WithErr(global.ConfigPathError ,c)
	controller.WithData(gin.H{
		"msg":"example success",
	},c)
}

func NewExampleController() *ExampleController{
	return &ExampleController{
		baseController: baseController{},
		exampleService: service.NewExampleService(),
	}
}