package router

import (
	"KokoChatting/controller"
	"github.com/gin-gonic/gin"
)

type userRouter struct{}

// 注册接口路由注册
func (r *userRouter) UserRouter(_router gin.IRoutes) {
	userController := controller.NewUserController()
	_router.POST("/user/register", userController.Register)
	_router.POST("/user/login", userController.Login)
}

func (r *userRouter) JWTUserRouter(_route gin.IRoutes) {
	userController := controller.NewUserController()
	_route.GET("/user", userController.GetUserInfo)
}