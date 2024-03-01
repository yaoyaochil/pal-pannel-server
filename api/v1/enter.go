package v1

import (
	"web-server/api/v1/pal_server"
	"web-server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup    system.ApiGroup
	PalServerApiGroup pal_server.PalServerOptionApi
}

var ApiGroupApp = new(ApiGroup)
