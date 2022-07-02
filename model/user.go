package model

import "time"

type User struct {
	BaseModal
	Mobile   string     `gorm:"index:index_mobile;unique;type:varchar(11);not null" json:"mobile"`
	Password string     `gorm:"type:varchar(100);not null" json:"password"`
	NickName string     `gorm:"type:varchar(20)" json:"nickName"`
	Birthday *time.Time `gorm:"type:datetime" json:"birthday"`
	Gender   string     `gorm:"type:varchar(12)" json:"gender"`
	Role     int        `gorm:"type:int;column:role;default:1"`
}
