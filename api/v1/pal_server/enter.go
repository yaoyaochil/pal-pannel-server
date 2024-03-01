package pal_server

import "web-server/service"

type ApiGroup struct {
	PalServerOptionApi
}

var (
	palOptionService = service.ServiceGroupApp.PalServerServiceGroup
)
