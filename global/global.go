package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"web-server/config"
)

var (
	PalDB                 *gorm.DB
	PalLog                *zap.Logger
	PalRedis              *redis.Client
	PalCache              local_cache.Cache
	PalConfig             config.Server
	PalVP                 *viper.Viper
	PalConcurrencyControl = &singleflight.Group{}
)
