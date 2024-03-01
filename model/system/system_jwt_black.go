package system

import "web-server/global"

type JwtBlacklist struct {
	global.Pal_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
