package initialize

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"web-server/global"
	"web-server/model/system"
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
	)
	if err != nil {
		global.PalLog.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.PalLog.Info("register table success")
}
