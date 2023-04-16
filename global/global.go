package global

import (
	"gorm.io/gorm"
	"sales-user-srv/model"
)

var (
	NacosInfo    *model.NacosConfig
	ServerConfig *model.ServerConfig
	DB           *gorm.DB
)
