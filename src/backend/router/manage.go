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
	_router.POST("/create_group", routerController.CreatGroup)
	_router.POST("/group/quit", routerController.QuitGroup)
	_router.GET("/user/list", routerController.GetFriendListInfo)
	_router.POST("/group/avatar", routerController.SetGroupAvatar)
	_router.GET("/group/list", routerController.GetGroupListInfo)
	_router.POST("/group/host/transfer", routerController.TransferHost)
	_router.POST("/group/host/change_permission", routerController.ChangePermission)
}
