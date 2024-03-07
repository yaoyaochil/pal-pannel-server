package pal_server

import (
	"web-server/model/common/response"
	"web-server/model/pal_server/request"

	globalRequest "web-server/model/common/request"

	"github.com/gin-gonic/gin"
)

type PalArchiveApi struct{}

// GetArchives 获取存档列表
func (a *PalArchiveApi) GetArchives(c *gin.Context) {
	var info request.ArchiveList
	_ = c.ShouldBindQuery(&info)
	data, total, err := palOptionService.PalSaveArchiveService.GetArchiveList(info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     data,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}, "获取成功", c)
}

// SaveArchive 存档保存
func (a *PalArchiveApi) SaveArchive(c *gin.Context) {
	var info request.ArchiveSave
	_ = c.ShouldBindJSON(&info)
	err := palOptionService.PalSaveArchiveService.SaveArchive(info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("备份成功", c)
}

// RestoreArchive 存档恢复
func (a *PalArchiveApi) RestoreArchive(c *gin.Context) {
	var info request.ArchiveRestore
	_ = c.ShouldBindJSON(&info)
	err := palOptionService.PalSaveArchiveService.RestoreArchive(info.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("恢复成功", c)
}

// DeleteArchive 存档删除
func (a *PalArchiveApi) DeleteArchive(c *gin.Context) {
	var info globalRequest.GetById
	_ = c.ShouldBindJSON(&info)
	err := palOptionService.PalSaveArchiveService.DeleteArchive(info.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
