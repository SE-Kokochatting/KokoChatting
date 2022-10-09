package router

import (
	"KokoChatting/controller"
	"github.com/gin-gonic/gin"
)

type upgradeRouter struct{}

func (r *upgradeRouter) RegisterUpgradeRouter(_router gin.IRoutes) {
	upgradeController := controller.NewUpgradeController()

	_router.GET("/upgrade_protocol",upgradeController.UpgradeProtocol)
}