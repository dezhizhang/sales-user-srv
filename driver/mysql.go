package driver

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sales-user-srv/model"
)

var DB *gorm.DB

func init() {
	dsn := "root:701XTAY1993@tcp(127.0.0.1:3306)/sales_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{})
	DB = db
}
