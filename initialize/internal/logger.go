package internal

import (
	"fmt"
	"web-server/global"

	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
// Author [wangrui19970405](https://github.com/wangrui19970405)
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
// Author [wangrui19970405](https://github.com/wangrui19970405)
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.PalConfig.System.DbType {
	case "mysql":
		logZap = global.PalConfig.Mysql.LogZap
	}
	if logZap {
		global.PalLog.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
