package router

import (
	"KokoChatting/controller"
	"github.com/gin-gonic/gin"
)

type userRouter struct{}

// 注册接口路由注册
func (r *userRouter) RegisterRoute(_router gin.IRouter) {
	routerController := controller.NewRegisterController()
	_router.POST("/user/register", routerController.Register)
}