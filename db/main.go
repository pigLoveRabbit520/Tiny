package db

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

var (
	DB        *gorm.DB
	ErrDBNull = errors.New("db is null, please connect first")
)

func ConnectAndInit(models ...interface{}) {
	dbUser := beego.AppConfig.String("dbUser")
	dbPass := beego.AppConfig.String("dbPass")
	dbHost := beego.AppConfig.String("dbHost")
	dbName := beego.AppConfig.String("dbName")
	dbPort := beego.AppConfig.String("dbPort")
	dbPrefix := beego.AppConfig.String("dbPrefix")

	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return dbPrefix + defaultTableName;
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
