package global

import (
	"gorm.io/gorm"
	"sales-user-srv/model"
)

var (
	ServerConfig *model.SalesServerConfig = &model.SalesServerConfig{}
	DB           *gorm.DB
)
