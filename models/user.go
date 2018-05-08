package models

import (
	"github.com/salamander-mh/SalamanderBlog/db"
	"time"
)

type User struct {
	Uid            int       `gorm:"primary_key"`
	Username       string    `gorm:"unique;type:varchar(30)"`
	Password       string    `gorm:"type:varchar(50)"`
	CreateDatetime time.Time
}

func GetUser(uid int) User {
	user := User{}
	db.DB.Where("uid = ?", "jinzhu").First(&user)
	return user
}
