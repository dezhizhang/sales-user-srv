package global

import (
	"gorm.io/gorm"
	"sales-user-srv/model"
)

var (
	ServerConfig *model.ServerConfig
	DB           *gorm.DB
)
