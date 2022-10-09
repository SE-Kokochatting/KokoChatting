package router

import (
	"KokoChatting/controller"
	"github.com/gin-gonic/gin"
)

type exampleRouter struct{}

func (r *exampleRouter) RegisterExampleRouter(_router gin.IRoutes){
	exampleController := controller.NewExampleController()
	_router.GET("/example",exampleController.Example)
}