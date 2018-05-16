package models

import (
	"github.com/salamander-mh/SalamanderBlog/db"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Uid        int    `gorm:"primary_key"  json:"uid"`
	Name       string `gorm:"type:varchar(32);unique" json:"name"`
	Password   string `gorm:"type:varchar(32)" json:"_"`
	Mail       string `gorm:"type:varchar(200)" json:"mail"`
	Url        string `gorm:"type:varchar(200)" json:"url"`
	ScreenName string `gorm:"type:varchar(32);column:screenName" json:"screenName"`
	Created    int    `gorm:"type:int(10) unsigned" json:"created"`
	Activated  int    `gorm:"type:int(10) unsigned" json:"activated"`
	Logged     int    `gorm:"type:int(10) unsigned" json:"logged"`
	Group      string `gorm:"type:varchar(16);default:'visitor'" json:"group"` // visitor || administrator
	AuthCode   string `gorm:"type:varchar(40)" json:"authCode"`
}

func GetUser(uid int) User {
	user := User{}
	db.DB.Where("uid = ?", uid).First(&user)
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
