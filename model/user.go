package model

type User struct {
	ID       string `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Mobile   string `gorm:"index:index_mobile;unique;type:varchar(11);not null" json:"mobile"`
	Password string `gorm:"type:varchar(100);not null" json:"password"`
	NickName string `gorm:"type:varchar(20)" json:"nickName"`
	Gender   string `gorm:"type: varchar(6); comment 'female 表示女,male 表示男'" json:"gender"`
	Role     int    `gorm:"type: int;column:role;default:1"`
}
