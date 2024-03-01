package core

import (
	"fmt"
	"go.uber.org/zap"
	"time"
	"web-server/global"
	"web-server/initialize"
	"web-server/service/system"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	// 初始化redis服务
	initialize.RedisInit()

	// 从db加载jwt数据
	if global.PalDB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.PalConfig.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.PalLog.Info("pal_server run success on ", zap.String("address", address))

	fmt.Printf(`
    帕鲁管理服务启动成功
	药要吃-私人播客:https://moonwife.top
`, address)
	global.PalLog.Error(s.ListenAndServe().Error())
}
