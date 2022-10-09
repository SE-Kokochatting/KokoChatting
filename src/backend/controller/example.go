package controller

import (
	"KokoChatting/service"
	"github.com/gin-gonic/gin"
)

type ExampleController struct{
	exampleService *service.ExampleService
}

func (controller *ExampleController) Example(c *gin.Context){
	err := controller.exampleService.Example(1,2,3,"test","example")
	if err != nil{
		c.JSON(200,gin.H{
			"msg":"error: "+err.Error(),
		})
	}
	c.JSON(200,gin.H{
		"msg":"hello world!!",
	})
}

func NewExampleController()*ExampleController{
	return &ExampleController{
		exampleService: service.NewExampleService(),
	}
}