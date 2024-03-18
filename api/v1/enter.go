package v1

import (
	"web-server/api/v1/pal_server"
	"web-server/api/v1/sokect"
	"web-server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup    system.ApiGroup
	PalServerApiGroup pal_server.ApiGroup
	SocketApiGroup    sokect.SocketApiGroup
}

var ApiGroupApp = new(ApiGroup)
