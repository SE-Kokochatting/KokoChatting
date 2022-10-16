package controller

import (
	"KokoChatting/global"
	"KokoChatting/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type upgradeController struct {
	baseController // 匿名字段（继承）
	upgrader       *websocket.Upgrader
	wsservice      *service.WsService
}

func (controller *upgradeController) UpgradeProtocol(c *gin.Context) {
	conn, err := controller.upgrader.Upgrade(c.Writer, c.Request, http.Header{})
	if err != nil {
		// 日志打印
		global.Logger.Error("conn upgrade error", zap.Error(err))
	}
	err = controller.wsservice.AddConn(conn)
	if err != nil {
		// 日志打印
		global.Logger.Error("add conn to ws conn managers error", zap.Error(err))
	}

}

func NewUpgradeController() *upgradeController {
	return &upgradeController{
		baseController: baseController{},
		upgrader:       new(websocket.Upgrader), // 暂无特殊配置    todo 完成配置文件解析，读取配置并填入upgrader中
		wsservice:      new(service.WsService),
	}
}
