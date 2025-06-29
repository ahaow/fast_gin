package global

import (
	"fast_gin/config"
	"fast_gin/core"

	"gorm.io/gorm"
)

const Version = "0.0.1"

var (
	Config *config.Config
	Log    *core.Logger
	DB     *gorm.DB
	// Redis  *redis.Client
)
