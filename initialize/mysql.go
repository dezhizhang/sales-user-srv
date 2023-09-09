package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sales-user-srv/global"
	"sales-user-srv/model"
)

func InitDB() {
	name := global.ServerConfig.MysqlConfig.Name
	host := global.ServerConfig.MysqlConfig.Host
	user := global.ServerConfig.MysqlConfig.User
	port := global.ServerConfig.MysqlConfig.Port
	password := global.ServerConfig.MysqlConfig.Password
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		user, password, host, port, name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{})
	global.DB = db
}
