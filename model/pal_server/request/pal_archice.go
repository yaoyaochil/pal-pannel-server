package request

import (
	"web-server/model/common/request"
	model "web-server/model/pal_server"
)

type ArchiveSave struct {
	ArchiveName string `json:"archiveName"`
	ArchiveDesc string `json:"archiveDesc"`
}

type ArchiveRestore struct {
	ID uint `json:"id"`
}

type ArchiveList struct {
	request.PageInfo
	model.PalArchive
}
