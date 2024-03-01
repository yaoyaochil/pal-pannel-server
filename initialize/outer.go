package initialize

import (
	"web-server/global"
	"web-server/utils"

	"github.com/songzhibin97/gkit/cache/local_cache"
)

func OtherInit() {
	dr, err := utils.ParseDuration(global.PalConfig.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.PalConfig.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	global.PalCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
}
