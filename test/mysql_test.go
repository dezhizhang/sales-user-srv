package test

import (
	"fmt"
	"log"
	"sales-user-srv/driver"
	"sales-user-srv/model"
	"testing"
)

func TestMySql(t *testing.T) {
	err := driver.DB.Find(&model.User{}).Error
	if err != nil {
		log.Printf("创建失败%s", err.Error())
	}

	fmt.Println("创建成功")
}
