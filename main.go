package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/salamander-mh/SalamanderBlog/db"
	"github.com/salamander-mh/SalamanderBlog/models"
)

func main() {
	db.ConnectAndInit(new(models.User), new(models.Content))
	beego.Run()
}