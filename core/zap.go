package core

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"web-server/core/internal"
	"web-server/global"
	"web-server/utils"
)

// Zap 获取 zap.Logger
// Author [wangrui19970405](https://github.com/wangrui19970405)
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.PalConfig.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.PalConfig.Zap.Director)
		_ = os.Mkdir(global.PalConfig.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.PalConfig.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
