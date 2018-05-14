package models

import (
	"github.com/salamander-mh/SalamanderBlog/db"
)

type User struct {
	Uid        int    `gorm:"primary_key"`
	Name       string `gorm:"type:varchar(32);unique"`
	Password   string `gorm:"type:varchar(32)"`
	Mail       string `gorm:"type:varchar(200)"`
	Url        string `gorm:"type:varchar(200)"`
	ScreenName string `gorm:"type:varchar(32);column:screenName"`
	Created    int    `gorm:"type:int(10) unsigned"`
	Activated  int    `gorm:"type:int(10) unsigned"`
	Logged     int    `gorm:"type:int(10) unsigned"`
	Group      string `gorm:"type:varchar(16);default:'visitor'"` // visitor || administrator
	AuthCode   string `gorm:"type:varchar(40)"`
}

func GetUser(uid int) User {
	user := User{}
	db.DB.Where("uid = ?", "jinzhu").First(&user)
	return user
}
