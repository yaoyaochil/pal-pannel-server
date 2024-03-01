package pal_server

import (
	"web-server/api/v1"
)

type RouterGroup struct {
	PalServerOptionRouter
}

var (
	palServerApi = v1.ApiGroupApp.PalServerApiGroup
)
