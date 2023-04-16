package driver

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sales-user-srv/global"
	"sales-user-srv/model"
)

func InitDB() (err error) {
	dsn := "root:701XTAY1993@tcp(127.0.0.1:3306)/sales?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println("dbs-----", dsn)
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		return err
	}
	return nil
}

func init() {
	InitDB()
	global.DB.AutoMigrate(&model.User{})
}
