package models

import (
	"github.com/salamander-mh/SalamanderBlog/db"
	"golang.org/x/crypto/bcrypt"
	"time"
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

func LoginUser(name, pass string) (*User, error) {
	user := User{}
	db.DB.Where("name = ?", name).Select("uid, password").First(&user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if err != nil {
		return nil, err
	}
	// 更新最后登录时间
	db.DB.Model(&user).Where("uid = ?", user.Uid).Update("logged", time.Now().Unix())
	return &user, nil
}
