package pal_server

import (
	"github.com/gin-gonic/gin"
)

type PalServerOptionRouter struct {
}

func (p *PalServerOptionRouter) InitPalServerOptionRouter(router *gin.RouterGroup) {
	r := router.Group("server")
	{
		r.GET("getOptions", palServerApi.GetOptions)
		r.POST("updateOptions", palServerApi.UpdateOptions)
		r.POST("unlockTo20", palServerApi.UnlockTo20)
		r.GET("checkUnlockTo20", palServerApi.CheckUnlockTo20)
	}
}
