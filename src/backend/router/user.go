package router

import (
	"KokoChatting/controller"
	"github.com/gin-gonic/gin"
)

type userRouter struct{}

// 注册接口路由注册
func (r *userRouter) UserRouter(_route gin.IRoutes) {
	userController := controller.NewUserController()
	_route.POST("/user/register", userController.Register)
	_route.POST("/user/login", userController.Login)
	_route.POST("/user", userController.GetUserInfo)
}

func (r *userRouter) JWTUserRouter(_route gin.IRoutes) {
	userController := controller.NewUserController()
	_route.POST("/user/avatar", userController.SetUserAvatar)
}