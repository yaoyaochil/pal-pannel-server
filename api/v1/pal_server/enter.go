package pal_server

import "web-server/service"

type ApiGroup struct {
	PalServerOptionApi
	PalArchiveApi
	PalInitApi
}

var (
	palOptionService = service.ServiceGroupApp.PalServerServiceGroup
)
