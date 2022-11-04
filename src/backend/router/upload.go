package router

import (
	"KokoChatting/controller"

	"github.com/gin-gonic/gin"
)

type uploadRouter struct{}

func (r *uploadRouter) UploadPictureRouter(_router gin.IRoutes) {
	uploadController := controller.NewUploadController()
	_router.POST("/upload",uploadController.Upload)
}