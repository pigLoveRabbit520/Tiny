package db

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	"errors"
)

var (
	DB *gorm.DB
	ErrDBNull = errors.New("db is null, please connect First")
)

func ConnectAndInit(models ...interface{}) {
	dbUser := beego.AppConfig.String("mysqluser")
	dbPass := beego.AppConfig.String("mysqlpass")
	dbHost := beego.AppConfig.String("mysqlhost")
	dbName := beego.AppConfig.String("mysqldb")
	dbPort := beego.AppConfig.String("mysqlport")
	Connect(dbUser, dbPass, dbHost, dbPort, dbName)
	DB.SingularTable(true)
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
