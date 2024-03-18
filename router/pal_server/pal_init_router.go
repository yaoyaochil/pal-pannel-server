package pal_server

import "github.com/gin-gonic/gin"

type PalInitRouter struct{}

// InitRouter 初始化pal路由
func (s *PalInitRouter) InitRouter(Router *gin.RouterGroup) {

	router := Router.Group("pal-init")
	{
		router.GET("checkInit", palServerApi.CheckInit)
		router.GET("logOut", palServerApi.LogOut)
	}
	{
		router.POST("init", palServerApi.Init)
	}
}
