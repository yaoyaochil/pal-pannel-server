package pal_server

import v1 "web-server/api/v1"

type RouterGroup struct {
	PalServerOptionRouter
	PalArchiveRouter
	PalInitRouter
}

var (
	palServerApi = v1.ApiGroupApp.PalServerApiGroup
)
