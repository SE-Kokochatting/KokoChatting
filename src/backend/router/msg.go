package router

import (
	"KokoChatting/controller"
	"github.com/gin-gonic/gin"
)

type msgRouter struct{}

func (r *msgRouter) RegisterMsgRouter(_route gin.IRoutes) {
	pushController := controller.NewPushController()
	_route.POST("/user/friend/send_message",pushController.SendMsg)
	_route.POST("/user/friend/revert_message",pushController.RevertMessage)
}