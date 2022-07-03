package model

import "time"

type User struct {
	BaseModal
	Name     string     `gorm:"type:varchar(100);not null" json:"name"`
	Mobile   string     `gorm:"index:index_mobile;unique;type:varchar(11);not null" json:"mobile"`
	Password string     `gorm:"type:varchar(100);not null" json:"password"`
	NickName string     `gorm:"type:varchar(20)" json:"nickName"`
	Birthday *time.Time `gorm:"type:datetime" json:"birthday"`
	Gender   int        `gorm:"type:int" json:"gender"`
	Role     int        `gorm:"type:int;column:role;default:1"`
}
