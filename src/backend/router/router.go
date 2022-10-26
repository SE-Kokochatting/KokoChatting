package router

import (
	"KokoChatting/controller"

	"github.com/gin-gonic/gin"
)

type routerGroup struct{
	upgrade upgradeRouter
	example exampleRouter
	user userRouter
	manage manageRouter
}


// placeholder 占位用函数，可删除
func dummyCode(r gin.IRoutes){}

func Routers() *gin.Engine {
	middleware := new(controller.Middleware)
	engine := gin.New()
	engine.Use(middleware.ZapLogger(), gin.Recovery(),middleware.CORS())  // 日志中间件使用zap替换gin原生日志库
	rg := new(routerGroup)

	privateGroup := engine.Group("/api/v1").Use(middleware.JwtAuthValidate()) // use方法的参数需得是jwt鉴权和cors等 中间件（请求拦截器）
	{
		rg.manage.ManageRouter(privateGroup)  // 需要加入登录鉴权的接口的router group注册时需要传入的routes命名为privateGroup
		rg.user.JWTUserRouter(privateGroup)
	}

	publicGroup := engine.Group("/api/v1")   // 无需use中间件的routes命名为publicGroup
	{
		dummyCode(publicGroup)  // 占位，无用代码可删除
		rg.example.RegisterExampleRouter(publicGroup)
		rg.user.UserRouter(publicGroup)
	}
	return engine
}