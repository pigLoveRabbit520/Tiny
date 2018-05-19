package models

import (
	"github.com/salamander-mh/SalamanderBlog/db"
	"golang.org/x/crypto/bcrypt"
	"time"
	"errors"
)

type User struct {
	Uid         uint   `gorm:"primary_key"  json:"uid"`
	Name        string `gorm:"type:varchar(32);unique" json:"name"`
	Password    string `gorm:"type:varchar(64)" json:"_"`
	PasswordNew string `gorm:"type:varchar(64)" json:"_"`
	Mail        string `gorm:"type:varchar(200)" json:"mail"`
	Url         string `gorm:"type:varchar(200)" json:"url"`
	ScreenName  string `gorm:"type:varchar(32);column:screenName" json:"screenName"`
	Created     uint   `gorm:"type:int(10) unsigned" json:"created"`
	Activated   uint   `gorm:"type:int(10) unsigned" json:"activated"`
	Logged      uint   `gorm:"type:int(10) unsigned" json:"logged"`
	Group       string `gorm:"type:varchar(16);default:'visitor'" json:"group"` // visitor || administrator
	AuthCode    string `gorm:"type:varchar(40)" json:"authCode"`
}

func GetUser(uid int) User {
	user := User{}
	db.DB.Where("uid = ?", uid).First(&user)
	return user
}

func LoginUser(name, pass string) (*User, error) {
	user := User{}
	db.DB.Where("name = ?", name).Select("uid, password_new").First(&user)
	if user.Uid <= 0 {
		return nil, errors.New("不存在该用户")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordNew), []byte(pass))
	if err != nil {
		return nil, err
	}
	// 更新最后登录时间
	db.DB.Model(&user).Where("uid = ?", user.Uid).Update("logged", time.Now().Unix())
	return &user, nil
}
