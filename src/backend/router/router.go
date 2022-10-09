package router

import "github.com/gin-gonic/gin"

type routerGroup struct{
	upgrade upgradeRouter
	example exampleRouter
}

// placeholder 占位用函数，可删除
func dummyCode(r gin.IRoutes){}

func Routers() *gin.Engine {
	engine := gin.Default()
	rg := new(routerGroup)

	privateGroup := engine.Group("").Use() // use方法的参数需得是jwt鉴权和cors等 中间件（请求拦截器）
	{
		rg.upgrade.RegisterUpgradeRouter(privateGroup)  // 需要加入登录鉴权的接口的router group注册时需要传入的routes命名为privateGroup
	}

	publicGroup := engine.Group("")   // 无需use中间件的routes命名为publicGroup
	{
		dummyCode(publicGroup)  // 占位，无用代码可删除
		rg.example.RegisterExampleRouter(publicGroup)
	}
	return engine
}