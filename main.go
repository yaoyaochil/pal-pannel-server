package main

import (
	"web-server/core"
	"web-server/global"
	"web-server/initialize"
	"web-server/source"
)

func main() {
	global.PalVP = core.Viper() // 初始化Viper
	initialize.OtherInit()      // 初始化其他
	global.PalLog = core.Zap()  // 初始化zap日志库
	global.PalDB = initialize.Gorm()
	if global.PalDB != nil {
		initialize.RegisterDBTables(global.PalDB) // 初始化表
		db, _ := global.PalDB.DB()
		defer db.Close()
	}

	source.InitSystemDB() // 初始化系统用户

	core.RunWindowsServer()
}
