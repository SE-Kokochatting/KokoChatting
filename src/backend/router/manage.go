package router

import (
	"KokoChatting/controller"
	"github.com/gin-gonic/gin"
)

type manageRouter struct {}

func(r *manageRouter) DeleteFriendRouter(_router gin.IRouter){
	routerController := controller.NewDeleteFriendController()
	_router.POST("/user/list_delete", routerController.DeleteFriend)
}
