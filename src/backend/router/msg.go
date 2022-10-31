package router

import (
	"KokoChatting/controller"
	"github.com/gin-gonic/gin"
)

type msgRouter struct{}

func (r *msgRouter) RegisterMsgRouter(_route gin.IRoutes) {
	pushController := controller.NewPushController()
	msgPullController := controller.NewMsgPullController()
	_route.POST("/user/friend/send_message",pushController.SendMsg)
	_route.POST("/user/friend/pull_msg_outline", msgPullController.MsgPullOutline)
	_route.POST("/user/friend/pull_message", msgPullController.MsgPull)
}