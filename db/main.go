package db

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/salamander-mh/SalamanderBlog/config"
)

var (
	DB        *gorm.DB
	ErrDBNull = errors.New("db is null, please connect first")
)

func ConnectAndInit(conf config.Config, models ...interface{}) {
	dbUser := conf.DB.User
	dbPass := conf.DB.Password
	dbHost := conf.DB.Host
	dbName := conf.DB.Name
	dbPort := fmt.Sprintf("%d", conf.DB.Port)
	dbPrefix := conf.DB.Prefix

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		if defaultTableName == "meta" {
			defaultTableName = "metas"
		}
		return dbPrefix + defaultTableName
	}
	Connect(dbUser, dbPass, dbHost, dbPort, dbName)
	Init(models...)
}

func Init(models ...interface{}) error {
	if DB == nil {
		return ErrDBNull
	}
	// DB.LogMode(true)
	DB.AutoMigrate(models...)
	return nil
}

func Connect(user, password, host string, dbPort, dbName string) {
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host,
			dbPort, dbName))
	if err != nil {
		panic("connect database failed " + err.Error())
	}
	DB = db
}
