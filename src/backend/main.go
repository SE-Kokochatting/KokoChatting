package main

import (
	"KokoChatting/global"
	"KokoChatting/router"
	"KokoChatting/wsserver"

	"go.uber.org/zap"
)



func main(){
	engine := router.Routers()
	port,err := global.GetGlobalConfig().GetConfigByPath("server.port")
	if err != nil{
		panic(err)
	}
	// start ws server
	wsserver.Run()

	if err := engine.Run("localhost:"+port); err != nil{
		global.Logger.Error("server run error",zap.Error(err))
	}
}