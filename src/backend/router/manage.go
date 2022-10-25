package router

import (
	"KokoChatting/controller"
	"github.com/gin-gonic/gin"
)

type manageRouter struct {}

func(r *manageRouter) ManageRouter(_router gin.IRoutes){
	routerController := controller.NewManageController()
	_router.POST("/user/list_delete", routerController.DeleteFriend)
	_router.POST("/user/list_block", routerController.BlockFriend)
	_router.POST("/creat_group", routerController.CreatGroup)
	_router.POST("/group/quit", routerController.QuitGroup)
}