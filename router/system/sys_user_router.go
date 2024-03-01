package system

import (
	v1 "web-server/api/v1"

	"github.com/gin-gonic/gin"
)

type SysUserRouter struct {
}

var sysUserApi = v1.ApiGroupApp.SystemApiGroup.SysUserApi

func (s *SysUserRouter) InitSysUserRouter(router *gin.RouterGroup) {
	Router := router.Group("user")
	{
		Router.GET("getUserInfo", sysUserApi.GetUserInfo)
	}
}
