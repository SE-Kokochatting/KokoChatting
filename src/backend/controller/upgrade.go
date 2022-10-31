package controller

import (
	"KokoChatting/global"
	"KokoChatting/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type upgradeController struct {
	baseController // 匿名字段（继承）
	upgrader       *websocket.Upgrader
	wsservice      *service.WsService
}

// todo: get uid from gin.Context
func (controller *upgradeController) UpgradeProtocol(c *gin.Context) {
	uid := controller.getUid(c)
	conn, err := controller.upgrader.Upgrade(c.Writer, c.Request, http.Header{})
	if err != nil {
		// 日志打印
		global.Logger.Error("conn upgrade error", zap.Error(err))
		controller.WithErr(global.UpgradeProtocolError,c)
		return
	}
	err = controller.wsservice.AddConn(conn,uid)
	if err != nil {
		// 日志打印
		global.Logger.Error("add conn to ws conn managers error", zap.Error(err))
		controller.WithData(struct{}{},c)
	}

}

func NewUpgradeController() *upgradeController {
	heartBeat, err := global.GetGlobalConfig().GetConfigByPath("server.websocket.health-check-duration")
	if err != nil {
		global.Logger.Error("server.websocket.health-check-duration get error", zap.Error(err))
		return nil
	}
	heartBeatDur, err := strconv.Atoi(heartBeat)
	return &upgradeController{
		baseController: baseController{},
		upgrader:       &websocket.Upgrader{
			HandshakeTimeout: time.Duration(heartBeatDur) * time.Second,
		}, // 暂无特殊配置   
		wsservice:      new(service.WsService),
	}
}
