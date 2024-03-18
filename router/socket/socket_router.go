package socket

import (
	"github.com/gin-gonic/gin"
	v1 "web-server/api/v1"
)

type SocketRouter struct{}

// InitSocketRouter 初始化socket路由
func (s *SocketRouter) InitSocketRouter(Router *gin.RouterGroup) {
	router := Router.Group("socket")
	socketApi := v1.ApiGroupApp.SocketApiGroup.SocketApi
	{
		router.GET("websocket", socketApi.Websocket)
	}
}
