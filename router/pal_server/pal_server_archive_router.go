package pal_server

import (
	v1 "web-server/api/v1"

	"github.com/gin-gonic/gin"
)

type PalArchiveRouter struct{}

// InitPalArchiveRouter 初始化存档路由
func (p *PalArchiveRouter) InitPalArchiveRouter(Router *gin.RouterGroup) {
	palArchiveApi := v1.ApiGroupApp.PalServerApiGroup.PalArchiveApi
	archive := Router.Group("archive")
	{
		archive.GET("getArchiveList", palArchiveApi.GetArchives)
		archive.POST("saveArchive", palArchiveApi.SaveArchive)
		archive.POST("restoreArchive", palArchiveApi.RestoreArchive)
		archive.POST("deleteArchive", palArchiveApi.DeleteArchive)
	}
}
