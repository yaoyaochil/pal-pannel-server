package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-server/middleware"
	"web-server/model/common/response"
	"web-server/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	// 自定义404
	Router.NoRoute(func(c *gin.Context) {
		response.NotFound(c)
	})

	//systemRouter := router.RouterGroupApp.System
	pal_server_router := router.RouterGroupApp.PalServer
	systemRouter := router.RouterGroupApp.System
	socketRouter := router.RouterGroupApp.Socket

	// 方便统一添加路由组前缀 多服务器上线使用
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth())
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		// 方便统一添加路由组前缀 多服务器上线使用
		//systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		pal_server_router.PalServerOptionRouter.InitPalServerOptionRouter(PublicGroup) // 注册幻兽帕鲁服务器配置路由
		pal_server_router.PalArchiveRouter.InitPalArchiveRouter(PublicGroup)           // 注册幻兽帕鲁存档路由
		pal_server_router.PalInitRouter.InitRouter(PublicGroup)                        // 注册幻兽帕鲁初始化路由
		systemRouter.SystemBaseRouter.InitSystemBaseRouter(PublicGroup)                // 注册用户路由
		systemRouter.SysJwtRouter.InitJwtRouter(PublicGroup)                           // 注册jwt相关路由

		socketRouter.SocketRouter.InitSocketRouter(PublicGroup) // 注册socket相关路由

	}
	{
		// 需要鉴权的路由
		systemRouter.SysUserRouter.InitSysUserRouter(PrivateGroup) // 注册用户路由
	}

	return Router
}
