package pal_server

import (
	"github.com/gin-gonic/gin"
	"web-server/model/common/response"
	model "web-server/model/pal_server"
)

type PalServerOptionApi struct{}

// GetOptions 获取幻兽帕鲁服务器配置
func (a *PalServerOptionApi) GetOptions(c *gin.Context) {
	data, err := palOptionService.Get()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(data, c)
}

// UpdateOptions 写入PalWorldSettings.ini文件
func (a *PalServerOptionApi) UpdateOptions(c *gin.Context) {
	var request model.PalServerOption
	if err := c.ShouldBind(&request); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := palOptionService.WriteIni(request)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("写入成功", c)
}
