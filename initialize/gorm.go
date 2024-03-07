package initialize

import (
	"os"
	"web-server/global"
	model "web-server/model/pal_server"
	"web-server/model/system"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm() *gorm.DB {
	switch global.PalConfig.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// RegisterDBTables 注册数据库表专用
func RegisterDBTables(db *gorm.DB) {
	err := db.AutoMigrate(
		// 系统模块表
		system.SysUser{},
		system.JwtBlacklist{},
		model.PalArchive{},
	)
	if err != nil {
		global.PalLog.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.PalLog.Info("register table success")
}
