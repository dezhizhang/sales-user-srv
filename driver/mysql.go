package driver

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sales-user-srv/model"
)

var DB *gorm.DB

func init() {
	var err error
	dsn := "root:701XTAY1993@tcp(127.0.0.1:3306)/sales_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	//global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		panic(err)
	}

	//global.DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.User{})
}
