package model

import "web-server/global"

type PalArchive struct {
	global.Pal_MODEL
	// 存档名称
	ArchiveName string `json:"archiveName" gorm:"comment:存档名称"`
	// 存档描述
	ArchiveDesc string `json:"archiveDesc" gorm:"comment:存档描述"`
	// 存档路径
	ArchivePath string `json:"archivePath" gorm:"comment:存档路径"`
	// 存档大小
	ArchiveSize int64 `json:"archiveSize" gorm:"comment:存档大小"`
}
