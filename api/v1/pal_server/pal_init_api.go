package pal_server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"web-server/global"
	"web-server/model/common/response"
	"web-server/utils"
)

type PalInitApi struct{}

// CheckInit 检查是否初始化
func (s *PalInitApi) CheckInit(c *gin.Context) {
	isInit := global.PalConfig.Palu.Init
	response.OkWithBool(isInit, "查询成功", c)
}

// Init 初始化
func (s *PalInitApi) Init(c *gin.Context) {
	scriptPath := "./init.sh"

	// 赋予执行权限
	err := os.Chmod(scriptPath, 0755)
	if err != nil {
		panic(err)
	}

	command := "./init.sh"
	go func() {
		err := utils.ExecuteCommandWithOutputToFile(command, global.PalConfig.Command.LogoutPath)
		if err != nil {
			global.PalLog.Error("初始化失败!", zap.Error(err))
			return
		}
	}()

	//global.PalConfig.Palu.Init = true

	response.OkWithMessage("初始化成功", c)
}

// LogOut 流式输出日志给前端
func (s *PalInitApi) LogOut(c *gin.Context) {
	filePath := "/home/steam/command.log"

	utils.StreamLogFile(filePath, c)
}
