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
	}
}
