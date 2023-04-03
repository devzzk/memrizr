package global

import (
	"github.com/devzzk/memrizr/config"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
)
