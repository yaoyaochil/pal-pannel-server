package router

import (
	"web-server/router/pal_server"
	"web-server/router/system"
)

type RouterGroup struct {
	System    system.RouterGroup
	PalServer pal_server.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
