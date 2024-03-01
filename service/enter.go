package service

import (
	"web-server/service/pal_server"
	"web-server/service/system"
)

type GroupService struct {
	SystemServiceGroup    system.ServiceGroup
	PalServerServiceGroup pal_server.ServiceGroup
}

var ServiceGroupApp = new(GroupService)
