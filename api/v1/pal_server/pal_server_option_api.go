package pal_server

import (
	"web-server/model/common/response"
	model "web-server/model/pal_server"

	"github.com/gin-gonic/gin"
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

type temType struct {
	IsUnlock bool `json:"isUnlock"`
}

// UnlockTo20 解锁据点上限到20
func (a *PalServerOptionApi) UnlockTo20(c *gin.Context) {
	var isunlock temType
	err := c.ShouldBindJSON(&isunlock)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = palOptionService.UnlockTo20(isunlock.IsUnlock)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("操作成功", c)
}

// checkUnlockTo20 解锁据点上限到20
func (a *PalServerOptionApi) CheckUnlockTo20(c *gin.Context) {
	isUnlock, err := palOptionService.CheckUnlockTo20()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	switch isUnlock {
	case true:
		response.OkWithBool(true, "据点上限解锁状态", c)
	case false:
		response.OkWithBool(false, "据点上限未解锁", c)
	}
}
